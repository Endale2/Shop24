// admin/controllers/seller_controller.go
package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/services"

	sellerModel "github.com/Endale2/DRPS/sellers/models"
	sellerServices "github.com/Endale2/DRPS/sellers/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateSeller handles creation of a new Seller (admin).
func CreateSeller(c *gin.Context) {
	var seller sellerModel.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := sellerServices.CreateSellerService(&seller)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating seller"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetSeller retrieves a single Seller by its ID, along with that seller's shops (id, name, slug, logo) and the seller's profile image.
func GetSeller(c *gin.Context) {
	id := c.Param("id")

	// 1) Fetch base seller record
	seller, err := sellerServices.GetSellerByIDService(id)
	if err != nil || seller == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}

	// 2) Fetch all shops and filter to this seller's shops
	allShops, err := services.GetAllShopsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving shops"})
		return
	}

	type ShopSummary struct {
		ID   primitive.ObjectID `json:"id"`
		Name string             `json:"name"`
		Slug string             `json:"slug"`
		Logo string             `json:"logo,omitempty"`
	}

	var ownedShops []ShopSummary
	for _, sh := range allShops {
		if sh.OwnerID == seller.ID {
			// Assuming the Shop model has a "Logo" field; if not, this will be empty
			ownedShops = append(ownedShops, ShopSummary{
				ID:   sh.ID,
				Name: sh.Name,
				Slug: sh.Slug,
				Logo: sh.Image, // or sh.Logo if defined in your Shop model
			})
		}
	}

	// 3) Build seller’s JSON including profile image
	//    (Assumes Seller model has a "ProfileImage" field)
	c.JSON(http.StatusOK, gin.H{
		"seller": gin.H{
			"id":            seller.ID,
			"first_name":    seller.FirstName,
			"last_name":     seller.LastName,
			"email":         seller.Email,
			"phone":         seller.Phone,
			"address":       seller.Address,
			"business_name": seller.BusinessName,
			"ratings":       seller.Ratings,
			"total_sales":   seller.TotalSales,
			"reviews_count": seller.ReviewsCount,
			"created_at":    seller.CreatedAt,
			"updated_at":    seller.UpdatedAt,
			"profile_image": seller.ProfileImage, // new field for seller’s avatar
		},
		"owned_shops": ownedShops,
	})
}

// GetSellers retrieves all Sellers (with only essential fields, not nested shops).
func GetSellers(c *gin.Context) {
	// 1) Fetch all sellers
	sellers, err := sellerServices.GetAllSellersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving sellers"})
		return
	}

	// 2) Build a slimmed-down list
	var response []gin.H
	for _, sl := range sellers {
		response = append(response, gin.H{
			"id":            sl.ID,
			"first_name":    sl.FirstName,
			"last_name":     sl.LastName,
			"business_name": sl.BusinessName,
			"email":         sl.Email,
			"ratings":       sl.Ratings,
			"profile_image": sl.ProfileImage, // include if set
		})
	}

	c.JSON(http.StatusOK, response)
}

// UpdateSeller updates an existing Seller by its ID.
func UpdateSeller(c *gin.Context) {
	id := c.Param("id")

	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := sellerServices.UpdateSellerService(id, bson.M(updatedData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating seller"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteSeller deletes a Seller by its ID.
func DeleteSeller(c *gin.Context) {
	id := c.Param("id")

	result, err := sellerServices.DeleteSellerService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting seller"})
		return
	}
	c.JSON(http.StatusOK, result)
}
