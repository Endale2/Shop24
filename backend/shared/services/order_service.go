// shared/services/order_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// generateOrderNumber creates a unique order number in format: ORD-YYYYMMDD-XXXXX
// where XXXXX is a sequential number for that day
func generateOrderNumber() (string, error) {
	today := time.Now().Format("20060102")

	// Get count of orders for today to generate sequential number
	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	filter := bson.M{
		"created_at": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	orders, err := repositories.ListOrders(context.Background(), filter)
	if err != nil {
		return "", err
	}

	// Generate sequential number (5 digits, zero-padded)
	sequence := len(orders) + 1
	orderNumber := fmt.Sprintf("ORD-%s-%05d", today, sequence)

	return orderNumber, nil
}

func CreateOrderService(o *models.Order) (*models.Order, error) {
	now := time.Now()
	o.CreatedAt = now
	o.UpdatedAt = now

	// Generate unique order number
	orderNumber, err := generateOrderNumber()
	if err != nil {
		return nil, err
	}
	o.OrderNumber = orderNumber

	res, err := repositories.CreateOrder(context.Background(), o)
	if err != nil {
		return nil, err
	}
	o.ID = res.InsertedID.(primitive.ObjectID)
	return o, nil
}

func GetOrderByIDService(idHex string) (*models.Order, error) {
	return repositories.GetOrderByID(context.Background(), idHex)
}

// GetOrderWithCustomerDetails returns order with customer information populated
func GetOrderWithCustomerDetails(idHex string) (*models.OrderWithCustomer, error) {
	order, err := repositories.GetOrderByID(context.Background(), idHex)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, nil
	}

	// Get customer details
	custModel, err := repositories.GetCustomerByID(order.CustomerID.Hex())
	if err != nil {
		return nil, err
	}

	// Convert customer model from customers package to shared models package
	customer := &models.Customer{
		ID:           custModel.ID,
		FirstName:    custModel.FirstName,
		LastName:     custModel.LastName,
		ProfileImage: custModel.ProfileImage,
		Email:        custModel.Email,
		Phone:        custModel.Phone,
		Address:      custModel.Address,
		City:         custModel.City,
		State:        custModel.State,
		PostalCode:   custModel.PostalCode,
		Country:      custModel.Country,
		CreatedAt:    custModel.CreatedAt,
		UpdatedAt:    custModel.UpdatedAt,
	}

	return &models.OrderWithCustomer{
		Order:    *order,
		Customer: customer,
	}, nil
}

func ListOrdersByShopService(shopIDHex string) ([]models.Order, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDHex)
	if err != nil {
		return nil, err
	}
	return repositories.ListOrders(context.Background(), bson.M{"shop_id": shopID})
}

// ListOrdersByShopPaginatedService returns paginated orders with customer information
func ListOrdersByShopPaginatedService(shopIDHex string, page, limit int, searchQuery string) ([]map[string]interface{}, int64, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDHex)
	if err != nil {
		return nil, 0, err
	}

	// Build filter
	filter := bson.M{"shop_id": shopID}

	// Add search filter if provided
	if searchQuery != "" {
		// Search in order number, customer ID, or product names
		searchRegex := primitive.Regex{Pattern: searchQuery, Options: "i"}
		filter["$or"] = []bson.M{
			{"order_number": searchRegex},
			{"customer_id": searchRegex},
			{"items.name": searchRegex},
		}
	}

	// Get paginated orders
	orders, total, err := repositories.ListOrdersPaginated(context.Background(), filter, page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Enhance orders with customer information
	var enhancedOrders []map[string]interface{}
	for _, order := range orders {
		orderMap := map[string]interface{}{
			"id":             order.ID,
			"order_number":   order.OrderNumber,
			"shop_id":        order.ShopID,
			"customer_id":    order.CustomerID,
			"items":          order.Items,
			"subtotal":       order.Subtotal,
			"discount_total": order.DiscountTotal,
			"total":          order.Total,
			"status":         order.Status,
			"created_at":     order.CreatedAt,
			"updated_at":     order.UpdatedAt,
		}

		// Try to get customer details
		customer, err := repositories.GetCustomerByID(order.CustomerID.Hex())
		if err != nil {
			// Set empty customer fields
			orderMap["customerFirstName"] = ""
			orderMap["customerLastName"] = ""
			orderMap["customerEmail"] = ""
		} else if customer != nil {
			orderMap["customerFirstName"] = customer.FirstName
			orderMap["customerLastName"] = customer.LastName
			orderMap["customerEmail"] = customer.Email
		} else {
			// Customer not found
			orderMap["customerFirstName"] = ""
			orderMap["customerLastName"] = ""
			orderMap["customerEmail"] = ""
		}

		enhancedOrders = append(enhancedOrders, orderMap)
	}

	return enhancedOrders, total, nil
}

// GetOrderStatsByShopService returns order statistics for a shop
func GetOrderStatsByShopService(shopIDHex string) (map[string]interface{}, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDHex)
	if err != nil {
		return nil, err
	}

	// Get all orders for the shop
	orders, err := repositories.ListOrders(context.Background(), bson.M{"shop_id": shopID})
	if err != nil {
		return nil, err
	}

	// Calculate statistics
	totalOrders := len(orders)
	totalRevenue := 0.0
	pendingOrders := 0
	deliveredOrders := 0

	for _, order := range orders {
		// Count orders by status
		switch order.Status {
		case "pending":
			pendingOrders++
		case "delivered":
			deliveredOrders++
		}

		// Calculate revenue for certain statuses (paid, shipped, delivered)
		if order.Status == "paid" || order.Status == "shipped" || order.Status == "delivered" {
			totalRevenue += order.Total
		}
	}

	return map[string]interface{}{
		"total_orders":     totalOrders,
		"total_revenue":    totalRevenue,
		"pending_orders":   pendingOrders,
		"delivered_orders": deliveredOrders,
	}, nil
}

func UpdateOrderService(idHex string, updates bson.M) (*models.Order, error) {
	updates["updated_at"] = time.Now()
	if _, err := repositories.UpdateOrder(context.Background(), idHex, updates); err != nil {
		return nil, err
	}
	return GetOrderByIDService(idHex)
}

func DeleteOrderService(idHex string) error {
	_, err := repositories.DeleteOrder(context.Background(), idHex)
	return err
}
