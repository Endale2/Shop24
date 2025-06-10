package services

import (
    "errors"
    "time"

    "github.com/Endale2/DRPS/shared/models"
    "github.com/Endale2/DRPS/shared/repositories"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrDiscountNotFound = errors.New("discount not found")

func CreateDiscountService(d *models.Discount) (*models.Discount, error) {
    if d.Name == "" {
        return nil, errors.New("discount name is required")
    }
    if d.ShopID.IsZero() {
        return nil, errors.New("shop ID is required")
    }
    if d.Category == "" {
        return nil, errors.New("discount category is required")
    }
    if !d.StartAt.IsZero() && !d.EndAt.IsZero() && d.EndAt.Before(d.StartAt) {
        return nil, errors.New("endAt must be after startAt")
    }
    // Optionally enforce coupon_code uniqueness: attempt repo.GetDiscountByCouponCode
    // ...
    res, err := repositories.CreateDiscount(d)
    if err != nil {
        return nil, err
    }
    if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
        d.ID = oid
    }
    return d, nil
}

func GetDiscountByIDService(idStr string) (*models.Discount, error) {
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        return nil, errors.New("invalid discount ID")
    }
    d, err := repositories.GetDiscountByID(id)
    if err != nil {
        return nil, err
    }
    if d == nil {
        return nil, ErrDiscountNotFound
    }
    return d, nil
}

func ListDiscountsByShopService(shopIDStr string) ([]models.Discount, error) {
    shopID, err := primitive.ObjectIDFromHex(shopIDStr)
    if err != nil {
        return nil, errors.New("invalid shop ID")
    }
    return repositories.ListDiscountsByShop(shopID)
}

func UpdateDiscountService(idStr string, upd bson.M) error {
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        return errors.New("invalid discount ID")
    }
    // remove immutable
    delete(upd, "id")
    delete(upd, "_id")
    if startRaw, ok := upd["start_at"]; ok {
        if startTime, ok2 := startRaw.(time.Time); ok2 {
            if endRaw, exists := upd["end_at"]; exists {
                if endTime, ok3 := endRaw.(time.Time); ok3 && endTime.Before(startTime) {
                    return errors.New("endAt must be after startAt")
                }
            }
        }
    }
    upd["updated_at"] = time.Now()
    _, err = repositories.UpdateDiscount(id, upd)
    return err
}

func DeleteDiscountService(idStr string) error {
    id, err := primitive.ObjectIDFromHex(idStr)
    if err != nil {
        return errors.New("invalid discount ID")
    }
    _, err = repositories.DeleteDiscount(id)
    return err
}

func GetDiscountByCouponCodeService(code string) (*models.Discount, error) {
    d, err := repositories.GetDiscountByCouponCode(code)
    if err != nil {
        return nil, err
    }
    if d == nil {
        return nil, ErrDiscountNotFound
    }
    now := time.Now()
    if !d.Active || now.Before(d.StartAt) || now.After(d.EndAt) {
        return nil, ErrDiscountNotFound
    }
    return d, nil
}

// For product-level: need collectionIDs slice if applicable
func GetActiveDiscountsForProductService(shopID, productID, variantID primitive.ObjectID, collectionIDs []primitive.ObjectID) ([]models.Discount, error) {
    return repositories.GetActiveDiscountsForProduct(shopID, productID, variantID, collectionIDs)
}

func GetActiveOrderDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
    return repositories.GetActiveOrderDiscountsForShop(shopID)
}

func GetActiveShippingDiscountsForShopService(shopID primitive.ObjectID) ([]models.Discount, error) {
    return repositories.GetActiveShippingDiscountsForShop(shopID)
}

func ValidateUsageLimits(discount *models.Discount, customerID primitive.ObjectID) (bool, error) {
    // stub or implement usage tracking
    return true, nil
}
