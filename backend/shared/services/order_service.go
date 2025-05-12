package services

import (
    "github.com/Endale2/DRPS/shared/models"
    "github.com/Endale2/DRPS/shared/repositories"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func CreateOrderService(order *models.Order) (*mongo.InsertOneResult, error) {
    // e.g. validate stock, calculate totals
    return repositories.CreateOrder(order)
}

func GetOrderByIDService(id string) (*models.Order, error) {
    return repositories.GetOrderByID(id)
}

func GetOrdersByUserService(userID string) ([]models.Order, error) {
    return repositories.GetOrdersByUser(userID)
}

func UpdateOrderStatusService(id string, statusUpdate bson.M) (*mongo.UpdateResult, error) {
    // e.g. validate status transitions
    return repositories.UpdateOrderStatus(id, statusUpdate)
}