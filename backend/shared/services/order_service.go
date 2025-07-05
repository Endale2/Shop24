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
