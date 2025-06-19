// shared/services/order_services.go
package services

import (
	"context"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrderService(o *models.Order) (*models.Order, error) {
	now := time.Now()
	o.CreatedAt = now
	o.UpdatedAt = now
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
