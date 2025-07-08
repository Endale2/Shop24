package controllers

import (
	"net/http"
	"sort"

	"github.com/Endale2/DRPS/shared/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTotalSales returns the total number of sales (orders)
func GetTotalSales(c *gin.Context) {
	count, err := repositories.CountOrders(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting sales"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_sales": count})
}

// GetTotalRevenue returns the total revenue from all orders
func GetTotalRevenue(c *gin.Context) {
	orders, err := repositories.ListOrders(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching orders"})
		return
	}
	total := 0.0
	for _, o := range orders {
		total += o.Total
	}
	c.JSON(http.StatusOK, gin.H{"total_revenue": total})
}

// GetOrderCount returns the total number of orders
func GetOrderCount(c *gin.Context) {
	count, err := repositories.CountOrders(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_count": count})
}

// GetTopProducts returns the top N products by sales count
func GetTopProducts(c *gin.Context) {
	orders, err := repositories.ListOrders(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching orders"})
		return
	}
	productSales := map[string]int{}
	for _, o := range orders {
		for _, item := range o.Items {
			productSales[item.ProductID.Hex()] += item.Quantity
		}
	}
	// Convert to slice and sort (top 5)
	type kv struct {
		Key   string
		Value int
	}
	var sorted []kv
	for k, v := range productSales {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Value > sorted[j].Value })
	if len(sorted) > 5 {
		sorted = sorted[:5]
	}
	c.JSON(http.StatusOK, sorted)
}

// GetTopCustomers returns the top N customers by order count
func GetTopCustomers(c *gin.Context) {
	orders, err := repositories.ListOrders(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching orders"})
		return
	}
	customerOrders := map[string]int{}
	for _, o := range orders {
		customerOrders[o.CustomerID.Hex()]++
	}
	type kv struct {
		Key   string
		Value int
	}
	var sorted []kv
	for k, v := range customerOrders {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Value > sorted[j].Value })
	if len(sorted) > 5 {
		sorted = sorted[:5]
	}
	c.JSON(http.StatusOK, sorted)
}
