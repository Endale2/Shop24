package services

import (
    "github.com/Endale2/DRPS/shared/models"
    "github.com/Endale2/DRPS/shared/repositories"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// LinkCustomerToShop creates the membership.
func LinkCustomerToShop(shopID, custID primitive.ObjectID) (*primitive.ObjectID, error) {
    link := &models.ShopCustomer{
        ShopID:     shopID,
        CustomerID: custID,
    }
    res, err := repositories.CreateShopCustomerLink(link)
    if err != nil {
        return nil, err
    }
    oid := res.InsertedID.(primitive.ObjectID)
    return &oid, nil
}

// Returns true if the given customerID is linked to shopID
func IsCustomerLinked(shopID, customerID primitive.ObjectID) (bool, error) {
    links, err := repositories.GetShopCustomerLinks(shopID)
    if err != nil {
        return false, err
    }
    for _, l := range links {
        if l.CustomerID == customerID {
            return true, nil
        }
    }
    return false, nil
}

// GetCustomerLinksForShop returns all ShopCustomer links for a shop.
func GetCustomerLinksForShop(shopID primitive.ObjectID) ([]models.ShopCustomer, error) {
    return repositories.GetShopCustomerLinks(shopID)
}

// UnlinkCustomerFromShop removes the link by its ID.
func UnlinkCustomerFromShop(linkID primitive.ObjectID) error {
    _, err := repositories.DeleteShopCustomerLink(linkID)
    return err
}
