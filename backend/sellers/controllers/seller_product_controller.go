package controllers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// variantInput represents a single variant with exactly one option.
type variantInput struct {
	OptionName  string  `json:"optionName" binding:"required"`
	OptionValue string  `json:"optionValue" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
}

// createProductInput represents the payload for creating a product.
type createProductInput struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Images      []string       `json:"images" binding:"required"`
	Category    string         `json:"category" binding:"required"`
	Price       float64        `json:"price" binding:"required"`
	CreatedBy   string         `json:"createdBy" binding:"required"`
	Variants    []variantInput `json:"variants"`
}

// slugify converts a name to a URL-friendly slug and ensures uniqueness within the shop.
func slugifyUnique(name string, shopID primitive.ObjectID) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	baseSlug := slug

	rand.Seed(time.Now().UnixNano())
	try := 0
	for {
		// Check if a product with this slug exists in this shop
		filter := bson.M{"slug": slug, "shop_id": shopID}
		products, _ := repositories.GetProductsByFilter(filter)
		if len(products) == 0 {
			return slug
		}
		// If exists, append random 4 chars
		slug = baseSlug + "-" + randSeq(4)
		try++
		if try > 5 { // fallback: add timestamp
			slug = baseSlug + "-" + time.Now().Format("20060102150405")
		}
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
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

	slug := slugifyUnique(in.Name, shopID)
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

	// Build variants: one option per variant
	for _, v := range in.Variants {
		p.Variants = append(p.Variants, models.Variant{
			Option: models.Option{
				Name:  v.OptionName,
				Value: v.OptionValue,
			},
			Price: v.Price,
			Stock: v.Stock,
			Image: v.Image,
		})
	}

	res, err := services.CreateProductService(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetProducts handles GET /seller/shops/:shopId/products
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

// GetProduct handles GET /seller/shops/:shopId/products/:productId
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

// UpdateProduct handles PATCH /seller/shops/:shopId/products/:productId
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

// DeleteProduct handles DELETE /seller/shops/:shopId/products/:productId
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
