package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	sellerRepositories "github.com/Endale2/DRPS/sellers/repositories"
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

// Helper: filter orders by status considered as 'sold'
func filterSoldOrders(orders []models.Order) []models.Order {
	var filtered []models.Order
	for _, o := range orders {
		status := o.Status
		if status == "paid" || status == "shipped" || status == "delivered" {
			filtered = append(filtered, o)
		}
	}
	return filtered
}

// GET /seller/shops/:shopId/analytics/summary
func GetShopAnalyticsSummary(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	orders = filterSoldOrders(orders)
	customers, _ := repositories.GetShopCustomerLinks(shopID)
	products, _ := services.GetProductsByShopIDService(shopID)

	totalRevenue := 0.0
	totalOrders := len(orders)
	customerOrderCount := make(map[primitive.ObjectID]int)
	for _, o := range orders {
		totalRevenue += o.Total
		customerOrderCount[o.CustomerID]++
	}
	totalCustomers := len(customers)
	newCustomers := 0
	returningCustomers := 0
	cutoff := time.Now().AddDate(0, 0, -30)
	for _, cst := range customers {
		if cst.CreatedAt.After(cutoff) {
			newCustomers++
		}
		if customerOrderCount[cst.CustomerID] > 1 {
			returningCustomers++
		}
	}
	avgOrderValue := 0.0
	if totalOrders > 0 {
		avgOrderValue = totalRevenue / float64(totalOrders)
	}

	c.JSON(http.StatusOK, gin.H{
		"total_revenue":       totalRevenue,
		"total_orders":        totalOrders,
		"new_customers":       newCustomers,
		"avg_order_value":     avgOrderValue,
		"total_products":      len(products),
		"total_customers":     totalCustomers,
		"returning_customers": returningCustomers,
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
	orders = filterSoldOrders(orders)

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
	orders = filterSoldOrders(orders)
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
	orders = filterSoldOrders(orders)
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
		collTitle := ""
		// Get the first collection title if the product has collections
		if len(p.CollectionIDs) > 0 {
			coll, _ := sellerRepositories.GetCollectionByID(p.CollectionIDs[0])
			if coll != nil {
				collTitle = coll.Title
			}
		}
		stats = append(stats, prodStat{
			ID:           p.ID,
			Name:         p.Name,
			Category:     collTitle,
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
	orders = filterSoldOrders(orders)
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

// GET /seller/shops/:shopId/analytics/customers-over-time?days=30
func GetShopCustomersOverTime(c *gin.Context) {
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
	customers, _ := repositories.GetShopCustomerLinks(shopID)
	growthMap := make(map[string]int)
	now := time.Now()
	for i := 0; i < days; i++ {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		growthMap[date] = 0
	}
	for _, cst := range customers {
		date := cst.CreatedAt.Format("2006-01-02")
		if _, ok := growthMap[date]; ok {
			growthMap[date]++
		}
	}
	var labels []string
	var values []int
	for i := days - 1; i >= 0; i-- {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		labels = append(labels, date)
		values = append(values, growthMap[date])
	}
	c.JSON(http.StatusOK, gin.H{"labels": labels, "values": values})
}

// GET /seller/shops/:shopId/analytics/category-sales?days=30
func GetShopCategorySales(c *gin.Context) {
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
	orders = filterSoldOrders(orders)
	products, _ := services.GetProductsByShopIDService(shopID)
	productCategory := make(map[primitive.ObjectID]string)
	for _, p := range products {
		collTitle := ""
		// Get the first collection title if the product has collections
		if len(p.CollectionIDs) > 0 {
			coll, _ := sellerRepositories.GetCollectionByID(p.CollectionIDs[0])
			if coll != nil {
				collTitle = coll.Title
			}
		}
		productCategory[p.ID] = collTitle
	}
	categoryRevenue := make(map[string]float64)
	now := time.Now()
	cutoff := now.AddDate(0, 0, -days)
	for _, o := range orders {
		if o.CreatedAt.Before(cutoff) {
			continue
		}
		for _, item := range o.Items {
			cat := productCategory[item.ProductID]
			if cat == "" {
				cat = "Uncategorized"
			}
			categoryRevenue[cat] += item.TotalPrice
		}
	}
	// Sort categories by revenue
	type catPair struct {
		Category string
		Revenue  float64
	}
	var pairs []catPair
	for cat, rev := range categoryRevenue {
		pairs = append(pairs, catPair{cat, rev})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Revenue > pairs[j].Revenue })
	if len(pairs) > 6 {
		pairs = pairs[:6]
	}
	var labels []string
	var values []float64
	for _, p := range pairs {
		labels = append(labels, p.Category)
		values = append(values, p.Revenue)
	}
	c.JSON(http.StatusOK, gin.H{"labels": labels, "values": values})
}

// GET /seller/shops/:shopId/analytics/recent-orders?limit=10
func GetShopRecentOrders(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}
	limit := 10
	if l := c.Query("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil {
			limit = n
		}
	}
	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	orders = filterSoldOrders(orders)
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.After(orders[j].CreatedAt)
	})
	if len(orders) > limit {
		orders = orders[:limit]
	}
	var result []gin.H
	for _, o := range orders {
		result = append(result, gin.H{
			"id":           o.ID.Hex(),
			"order_number": o.OrderNumber,
			"customer_id":  o.CustomerID.Hex(),
			"total":        o.Total,
			"status":       o.Status,
			"created_at":   o.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, result)
}

// GET /seller/shops/:shopId/analytics/dashboard
func GetShopDashboardAnalytics(c *gin.Context) {
	_, shopID, ok := getShopAndVerifySeller(c)
	if !ok {
		return
	}

	orders, _ := services.ListOrdersByShopService(shopID.Hex())
	products, _ := services.GetProductsByShopIDService(shopID)
	customers, _ := repositories.GetShopCustomerLinks(shopID)
	discounts, _ := repositories.ListDiscountsByShop(shopID)

	soldOrders := filterSoldOrders(orders)
	totalOrders := len(soldOrders)
	pendingOrders := 0
	newOrdersToday := 0
	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	totalRevenue := 0.0
	revenueToday := 0.0
	for _, o := range orders {
		if o.Status == "pending" {
			pendingOrders++
		}
		if o.CreatedAt.After(today) {
			if o.Status == "paid" || o.Status == "shipped" || o.Status == "delivered" {
				newOrdersToday++
				revenueToday += o.Total
			}
		}
	}
	for _, o := range soldOrders {
		totalRevenue += o.Total
	}
	avgOrderValue := 0.0
	if totalOrders > 0 {
		avgOrderValue = totalRevenue / float64(totalOrders)
	}

	totalProducts := len(products)
	lowStockProducts := 0
	for _, p := range products {
		if p.Stock < 10 {
			lowStockProducts++
		}
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i].CreatedAt.After(products[j].CreatedAt)
	})
	recentProducts := []gin.H{}
	maxRecent := 5
	if len(products) < maxRecent {
		maxRecent = len(products)
	}
	for i := 0; i < maxRecent; i++ {
		p := products[i]
		price := p.Price
		if len(p.Variants) > 0 {
			minPrice := p.Variants[0].Price
			for _, v := range p.Variants {
				if v.Price < minPrice {
					minPrice = v.Price
				}
			}
			price = minPrice
		}
		collTitle := ""
		// Get the first collection title if the product has collections
		if len(p.CollectionIDs) > 0 {
			coll, _ := sellerRepositories.GetCollectionByID(p.CollectionIDs[0])
			if coll != nil {
				collTitle = coll.Title
			}
		}
		recentProducts = append(recentProducts, gin.H{
			"id":        p.ID.Hex(),
			"name":      p.Name,
			"category":  collTitle,
			"mainImage": p.MainImage,
			"stock":     p.Stock,
			"createdAt": p.CreatedAt,
			"price":     price,
		})
	}
	totalCustomers := len(customers)
	newCustomersToday := 0
	for _, cst := range customers {
		if cst.CreatedAt.After(today) {
			newCustomersToday++
		}
	}
	activeDiscounts := 0
	for _, d := range discounts {
		if d.IsActive() {
			activeDiscounts++
		}
	}
	productSales := map[primitive.ObjectID]int{}
	for _, o := range soldOrders {
		for _, item := range o.Items {
			productSales[item.ProductID] += item.Quantity
		}
	}
	topProductID := primitive.NilObjectID
	topProductQty := 0
	for pid, qty := range productSales {
		if qty > topProductQty {
			topProductID = pid
			topProductQty = qty
		}
	}
	var topProduct *models.Product
	for _, p := range products {
		if p.ID == topProductID {
			topProduct = &p
			break
		}
	}
	sort.Slice(soldOrders, func(i, j int) bool {
		return soldOrders[i].CreatedAt.After(soldOrders[j].CreatedAt)
	})
	recentOrders := []models.Order{}
	if len(soldOrders) > 5 {
		recentOrders = soldOrders[:5]
	} else {
		recentOrders = soldOrders
	}

	// --- Weekly sales (last 7 days, by date) ---
	days := 7
	now := time.Now()
	revenueMap := make(map[string]float64)
	for i := 0; i < days; i++ {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		revenueMap[date] = 0
	}
	for _, o := range soldOrders {
		date := o.CreatedAt.Format("2006-01-02")
		if _, ok := revenueMap[date]; ok {
			revenueMap[date] += o.Total
		}
	}
	weeklySales := make([]gin.H, 0, days)
	weekdays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	for i := days - 1; i >= 0; i-- {
		date := now.AddDate(0, 0, -i)
		dateStr := date.Format("2006-01-02")
		weekday := weekdays[date.Weekday()]
		weeklySales = append(weeklySales, gin.H{
			"day":    weekday,
			"amount": revenueMap[dateStr],
		})
	}

	resp := gin.H{
		"total_products":      totalProducts,
		"low_stock_products":  lowStockProducts,
		"total_orders":        totalOrders,
		"pending_orders":      pendingOrders,
		"new_orders_today":    newOrdersToday,
		"total_customers":     totalCustomers,
		"new_customers_today": newCustomersToday,
		"total_revenue":       totalRevenue,
		"revenue_today":       revenueToday,
		"average_order_value": avgOrderValue,
		"active_discounts":    activeDiscounts,
		"recent_orders":       recentOrders,
		"recent_products":     recentProducts,
		"weekly_sales":        weeklySales,
	}
	if topProduct != nil {
		collTitle := ""
		// Get the first collection title if the product has collections
		if len(topProduct.CollectionIDs) > 0 {
			coll, _ := sellerRepositories.GetCollectionByID(topProduct.CollectionIDs[0])
			if coll != nil {
				collTitle = coll.Title
			}
		}
		resp["top_product"] = gin.H{
			"id":        topProduct.ID.Hex(),
			"name":      topProduct.Name,
			"category":  collTitle,
			"mainImage": topProduct.MainImage,
			"totalSold": topProductQty,
		}
	} else {
		resp["top_product"] = nil
	}

	c.JSON(http.StatusOK, resp)
}
