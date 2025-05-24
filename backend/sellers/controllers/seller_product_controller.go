package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// variantInput and createProductInput remain unchanged

type variantInput struct {
	Name    string    `json:"name" binding:"required"`
	Options []string  `json:"options" binding:"required"`
	Prices  []float64 `json:"prices" binding:"required"`
}

type createProductInput struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Images      []string       `json:"images" binding:"required"`
	Category    string         `json:"category" binding:"required"`
	Price       float64        `json:"price" binding:"required"`
	CreatedBy   string         `json:"createdBy" binding:"required"`
	Variants    []variantInput `json:"variants"`
}

// CreateProduct handles POST /seller/shops/:shopId/products
func CreateProduct(c *gin.Context) {
	// seller auth
	sellerHex, _ := c.Get("user_id")
	sellerID, err := primitive.ObjectIDFromHex(sellerHex.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

	// parse & authorize shop (route uses :shopId)
	shopParam := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	// bind payload
	var in createProductInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	// map to model
	p := &models.Product{
		ShopID:      shopID,
		UserID:      sellerID,
		Name:        in.Name,
		Description: in.Description,
		Images:      in.Images,
		Category:    in.Category,
		Price:       in.Price,
		CreatedBy:   sellerID,
	}

	// build variants
	for _, v := range in.Variants {
		if len(v.Options) != len(v.Prices) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "variant '" + v.Name + "' options/prices mismatch"})
			return
		}
		for i, opt := range v.Options {
			p.Variants = append(p.Variants, models.Variant{
				Options: map[string]string{v.Name: opt},
				Price:   v.Prices[i],
				Stock:   0,
			})
		}
	}

	// persist
	res, err := services.CreateProductService(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetProducts lists products for a seller’s shop.
// GET /seller/shops/:shopId/products
func GetProducts(c *gin.Context) {
	// seller auth
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

	// parse & authorize shop
	shopParam := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	// fetch filtered products via service
	products, err := services.GetProductsByShopIDService(shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct retrieves one product in a seller’s shop.
// GET /seller/shops/:shopId/products/:productId
func GetProduct(c *gin.Context) {
	// seller auth + shop authorize
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopParam := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	// load and verify product
	prID := c.Param("productId")
	p, err := services.GetProductByIDService(prID)
	if err != nil || p.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

// UpdateProduct updates a product in the seller’s shop.
// PATCH /seller/shops/:shopId/products/:productId
func UpdateProduct(c *gin.Context) {
	// auth + authorize
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopParam := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	// bind updates
	var upd bson.M
	if err := c.ShouldBindJSON(&upd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	upd["shop_id"] = shop.ID

	res, err := services.UpdateProductService(c.Param("productId"), upd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteProduct deletes a product from a seller’s shop.
// DELETE /seller/shops/:shopId/products/:productId
func DeleteProduct(c *gin.Context) {
	// auth + authorize
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopParam := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	res, err := services.DeleteProductService(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, res)
}
