package services

import (
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOrderService creates a new order.
func CreateOrderService(order *models.Order) (*mongo.InsertOneResult, error) {
	// Place for business validations (e.g., inventory check, price calculation)
	return repositories.CreateOrder(order)
}

// GetOrderByIDService retrieves a single order by its ID.
func GetOrderByIDService(id string) (*models.Order, error) {
	return repositories.GetOrderByID(id)
}

// GetAllOrdersService returns all orders.
func GetAllOrdersService() ([]models.Order, error) {
	return repositories.GetAllOrders()
}

// UpdateOrderService updates fields of an order identified by its ID.
func UpdateOrderService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	// Place for business rules (e.g., status transition validation)
	return repositories.UpdateOrder(id, updatedData)
}

// DeleteOrderService removes an order by its ID.
func DeleteOrderService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteOrder(id)
}
