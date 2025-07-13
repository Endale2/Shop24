package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// getShopAndVerifySeller ensures the shop exists and is owned by the current seller
func getShopAndVerifySeller(c *gin.Context) (*models.Shop, primitive.ObjectID, bool) {
	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return nil, primitive.NilObjectID, false
	}
	shop, err := services.GetShopByIDService(shopHex)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return nil, primitive.NilObjectID, false
	}
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	if shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized for this shop"})
		return nil, primitive.NilObjectID, false
	}
	return shop, shopID, true
}

// GET /seller/shops/:shopId/analytics/summary
func GetShopAnalyticsSummary(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	customers, _ := repositories.GetShopCustomerLinks(shopID)
	products, _ := services.GetProductsByShopIDService(shopID)

	totalRevenue := 0.0
	for _, o := range orders {
		totalRevenue += o.Total
	}
	totalOrders := len(orders)
	newCustomers := 0
	cutoff := time.Now().AddDate(0, 0, -30)
	for _, cst := range customers {
		if cst.CreatedAt.After(cutoff) {
			newCustomers++
		}
	}
	avgOrderValue := 0.0
	if totalOrders > 0 {
		avgOrderValue = totalRevenue / float64(totalOrders)
	}

	c.JSON(http.StatusOK, gin.H{
		"total_revenue":   totalRevenue,
		"total_orders":    totalOrders,
		"new_customers":   newCustomers,
		"avg_order_value": avgOrderValue,
		"total_products":  len(products),
	})
}

// GET /seller/shops/:shopId/analytics/revenue-over-time?days=30
func GetShopRevenueOverTime(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	days := 30
	if d := c.Query("days"); d != "" {
		if n, err := strconv.Atoi(d); err == nil {
			days = n
		}
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())

	// Build daily revenue
	revenueMap := make(map[string]float64)
	now := time.Now()
	for i := 0; i < days; i++ {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		revenueMap[date] = 0
	}
	for _, o := range orders {
		date := o.CreatedAt.Format("2006-01-02")
		if _, ok := revenueMap[date]; ok {
			revenueMap[date] += o.Total
		}
	}
	// Sort by date ascending
	var labels []string
	var values []float64
	for i := days - 1; i >= 0; i-- {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		labels = append(labels, date)
		values = append(values, revenueMap[date])
	}
	c.JSON(http.StatusOK, gin.H{"labels": labels, "values": values})
}

// GET /seller/shops/:shopId/analytics/orders-over-time?days=30
func GetShopOrdersOverTime(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	days := 30
	if d := c.Query("days"); d != "" {
		if n, err := strconv.Atoi(d); err == nil {
			days = n
		}
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	orderMap := make(map[string]int)
	now := time.Now()
	for i := 0; i < days; i++ {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		orderMap[date] = 0
	}
	for _, o := range orders {
		date := o.CreatedAt.Format("2006-01-02")
		if _, ok := orderMap[date]; ok {
			orderMap[date]++
		}
	}
	var labels []string
	var values []int
	for i := days - 1; i >= 0; i-- {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		labels = append(labels, date)
		values = append(values, orderMap[date])
	}
	c.JSON(http.StatusOK, gin.H{"labels": labels, "values": values})
}

// GET /seller/shops/:shopId/analytics/top-products
func GetShopTopProducts(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	products, _ := services.GetProductsByShopIDService(shopID)
	productSales := map[primitive.ObjectID]int{}
	productRevenue := map[primitive.ObjectID]float64{}
	for _, o := range orders {
		for _, item := range o.Items {
			productSales[item.ProductID] += item.Quantity
			productRevenue[item.ProductID] += item.TotalPrice
		}
	}

	type prodStat struct {
		ID           primitive.ObjectID `json:"id"`
		Name         string             `json:"name"`
		Category     string             `json:"category"`
		MainImage    string             `json:"mainImage"`
		TotalSold    int                `json:"totalSold"`
		TotalRevenue float64            `json:"totalRevenue"`
	}
	var stats []prodStat
	for _, p := range products {
		stats = append(stats, prodStat{
			ID:           p.ID,
			Name:         p.Name,
			Category:     p.Category,
			MainImage:    p.MainImage,
			TotalSold:    productSales[p.ID],
			TotalRevenue: productRevenue[p.ID],
		})
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].TotalRevenue > stats[j].TotalRevenue })
	if len(stats) > 5 {
		stats = stats[:5]
	}
	c.JSON(http.StatusOK, stats)
}

// GET /seller/shops/:shopId/analytics/top-customers
func GetShopTopCustomers(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	customerOrders := map[primitive.ObjectID]int{}
	for _, o := range orders {
		customerOrders[o.CustomerID]++
	}

	type custStat struct {
		ID    primitive.ObjectID `json:"id"`
		Count int                `json:"orderCount"`
	}
	var stats []custStat
	for cid, count := range customerOrders {
		stats = append(stats, custStat{ID: cid, Count: count})
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].Count > stats[j].Count })
	if len(stats) > 5 {
		stats = stats[:5]
	}
	c.JSON(http.StatusOK, stats)
}

// GET /seller/shops/:shopId/analytics/inventory-status
func GetShopInventoryStatus(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	products, _ := services.GetProductsByShopIDService(shopID)
	total := len(products)
	inStock := 0
	lowStock := 0
	outOfStock := 0
	for _, p := range products {
		if p.Stock > 10 {
			inStock++
		} else if p.Stock > 0 {
			lowStock++
		} else {
			outOfStock++
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"total_products": total,
		"in_stock":       inStock,
		"low_stock":      lowStock,
		"out_of_stock":   outOfStock,
	})
}

// GET /seller/shops/:shopId/analytics/discount-performance
func GetShopDiscountPerformance(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	discounts, _ := repositories.ListDiscountsByShop(shopID)

	type discStat struct {
		ID          primitive.ObjectID `json:"id"`
		Name        string             `json:"name"`
		TotalUsage  int                `json:"totalUsage"`
		TotalAmount float64            `json:"totalAmount"`
	}
	var stats []discStat
	for _, d := range discounts {
		amount := 0.0
		for _, usage := range d.UsageTracking {
			amount += usage.TotalSpent
		}
		stats = append(stats, discStat{
			ID:          d.ID,
			Name:        d.Name,
			TotalUsage:  d.CurrentUsage,
			TotalAmount: amount,
		})
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].TotalAmount > stats[j].TotalAmount })
	c.JSON(http.StatusOK, stats)
}
