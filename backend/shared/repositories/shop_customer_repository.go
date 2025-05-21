package repositories

import (
    "context"
    "errors"
    "time"

    "github.com/Endale2/DRPS/shared/models"
    "github.com/Endale2/DRPS/config"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var shopCustomerColl *mongo.Collection = config.GetCollection("DRPS", "shop_customers")

// CreateLink adds a Customer to a Shop.
func CreateShopCustomerLink(link *models.ShopCustomer) (*mongo.InsertOneResult, error) {
    link.ID = primitive.NewObjectID()
    link.CreatedAt = time.Now()
    link.UpdatedAt = time.Now()
    return shopCustomerColl.InsertOne(context.Background(), link)
}

// GetCustomersByShop returns all ShopCustomer links for one shop.
func GetShopCustomerLinks(shopID primitive.ObjectID) ([]models.ShopCustomer, error) {
    cursor, err := shopCustomerColl.Find(context.Background(), bson.M{"shopId": shopID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var links []models.ShopCustomer
    for cursor.Next(context.Background()) {
        var l models.ShopCustomer
        if err := cursor.Decode(&l); err != nil {
            return nil, err
        }
        links = append(links, l)
    }
    return links, nil
}

// DeleteLink removes a customer-shop association.
func DeleteShopCustomerLink(id primitive.ObjectID) (*mongo.DeleteResult, error) {
    if id.IsZero() {
        return nil, errors.New("invalid link ID")
    }
    return shopCustomerColl.DeleteOne(context.Background(), bson.M{"_id": id})
}
