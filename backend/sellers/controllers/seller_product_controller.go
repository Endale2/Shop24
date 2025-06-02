package controllers

import (
	"net/http"
	"strings"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

// slugify converts a name to a URL-friendly slug
func slugify(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	return slug
}

// CreateProduct handles POST /seller/shops/:shopId/products
func CreateProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, err := primitive.ObjectIDFromHex(sellerHex.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
		return
	}

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

	var in createProductInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	slug := slugify(in.Name)
	p := &models.Product{
		ShopID:      shopID,
		UserID:      sellerID,
		Name:        in.Name,
		Slug:        slug,
		Description: in.Description,
		Images:      in.Images,
		Category:    in.Category,
		Price:       in.Price,
		CreatedBy:   sellerID,
	}

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

	res, err := services.CreateProductService(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func GetProducts(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))

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

	products, err := services.GetProductsByShopIDService(shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fetch failed"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopParam := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	prID := c.Param("productId")
	p, err := services.GetProductByIDService(prID)
	if err != nil || p.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func UpdateProduct(c *gin.Context) {
	sellerHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(sellerHex.(string))
	shopParam := c.Param("shopId")
	shop, err := services.GetShopByIDService(shopParam)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

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

func DeleteProduct(c *gin.Context) {
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
