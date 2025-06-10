package repositories

import (
    "context"
    "time"

    "github.com/Endale2/DRPS/config"
    "github.com/Endale2/DRPS/shared/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var discountColl *mongo.Collection = config.GetCollection("DRPS", "discounts")

func CreateDiscount(d *models.Discount) (*mongo.InsertOneResult, error) {
    // ID and timestamps stamped by service or repo; choose one. If here:
    if d.ID.IsZero() {
        d.ID = primitive.NewObjectID()
    }
    now := time.Now()
    if d.CreatedAt.IsZero() {
        d.CreatedAt = now
    }
    d.UpdatedAt = now
    return discountColl.InsertOne(context.Background(), d)
}

func GetDiscountByID(id primitive.ObjectID) (*models.Discount, error) {
    var d models.Discount
    err := discountColl.FindOne(context.Background(), bson.M{"_id": id}).Decode(&d)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &d, nil
}

func ListDiscountsByShop(shopID primitive.ObjectID) ([]models.Discount, error) {
    cursor, err := discountColl.Find(context.Background(), bson.M{"shop_id": shopID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    var list []models.Discount
    for cursor.Next(context.Background()) {
        var d models.Discount
        if err := cursor.Decode(&d); err != nil {
            return nil, err
        }
        list = append(list, d)
    }
    return list, nil
}

func UpdateDiscount(id primitive.ObjectID, upd bson.M) (*mongo.UpdateResult, error) {
    upd["updated_at"] = time.Now()
    return discountColl.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": upd})
}

func DeleteDiscount(id primitive.ObjectID) (*mongo.DeleteResult, error) {
    return discountColl.DeleteOne(context.Background(), bson.M{"_id": id})
}

func GetDiscountByCouponCode(code string) (*models.Discount, error) {
    var d models.Discount
    err := discountColl.FindOne(context.Background(), bson.M{"coupon_code": code}).Decode(&d)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &d, nil
}

// Active queries: product-level including collections
func GetActiveDiscountsForProduct(shopID, productID, variantID primitive.ObjectID, collectionIDs []primitive.ObjectID) ([]models.Discount, error) {
    now := time.Now()
    orClauses := []bson.M{
        {"applies_to_products": productID},
        {"applies_to_variants": variantID},
    }
    if len(collectionIDs) > 0 {
        orClauses = append(orClauses, bson.M{"applies_to_collections": bson.M{"$in": collectionIDs}})
    }
    filter := bson.M{
        "shop_id":  shopID,
        "active":   true,
        "start_at": bson.M{"$lte": now},
        "end_at":   bson.M{"$gte": now},
        "category": models.DiscountCategoryProduct,
        "$or":      orClauses,
    }
    cursor, err := discountColl.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    var results []models.Discount
    for cursor.Next(context.Background()) {
        var d models.Discount
        if err := cursor.Decode(&d); err != nil {
            return nil, err
        }
        // Optionally check AllowedCustomerIDs / usage limits here or in service
        results = append(results, d)
    }
    return results, nil
}

func GetActiveOrderDiscountsForShop(shopID primitive.ObjectID) ([]models.Discount, error) {
    now := time.Now()
    filter := bson.M{
        "shop_id":  shopID,
        "active":   true,
        "start_at": bson.M{"$lte": now},
        "end_at":   bson.M{"$gte": now},
        "category": models.DiscountCategoryOrder,
    }
    cursor, err := discountColl.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    var list []models.Discount
    for cursor.Next(context.Background()) {
        var d models.Discount
        if err := cursor.Decode(&d); err != nil {
            return nil, err
        }
        list = append(list, d)
    }
    return list, nil
}

func GetActiveShippingDiscountsForShop(shopID primitive.ObjectID) ([]models.Discount, error) {
    now := time.Now()
    filter := bson.M{
        "shop_id":  shopID,
        "active":   true,
        "start_at": bson.M{"$lte": now},
        "end_at":   bson.M{"$gte": now},
        "category": models.DiscountCategoryShipping,
    }
    cursor, err := discountColl.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    var list []models.Discount
    for cursor.Next(context.Background()) {
        var d models.Discount
        if err := cursor.Decode(&d); err != nil {
            return nil, err
        }
        list = append(list, d)
    }
    return list, nil
}
