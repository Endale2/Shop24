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

// optionInput represents a single option for a variant.
type optionInput struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

// variantInput represents a single variant with multiple options.
type variantInput struct {
	Options []optionInput `json:"options" binding:"required"`
	Price   float64       `json:"price" binding:"required"`
	Stock   int           `json:"stock"`
	Image   string        `json:"image"`
}

// createProductInput represents the payload for creating a product.
type createProductInput struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description" binding:"required"`
	MainImage   string         `json:"main_image"` // <-- Added
	Images      []string       `json:"images" binding:"required"`
	Category    string         `json:"category" binding:"required"`
	Price       *float64       `json:"price"`
	Stock       *int           `json:"stock"` // <-- Added
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

	// Custom validation for price/variants
	if len(in.Variants) == 0 {
		if in.Price == nil || *in.Price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product price is required and must be positive if no variants are provided"})
			return
		}
	}

	// Build variants: multiple options per variant
	if len(in.Variants) > 0 {
		for _, v := range in.Variants {
			if len(v.Options) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Each variant must have at least one option (name:value)"})
				return
			}
			for _, o := range v.Options {
				if strings.TrimSpace(o.Name) == "" || strings.TrimSpace(o.Value) == "" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Variant option name and value cannot be empty"})
					return
				}
			}
		}
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
		Price:       0,
		CreatedBy:   sellerID,
	}

	p.MainImage = in.MainImage // <-- Set main_image
	if in.Stock != nil {
		p.Stock = *in.Stock // <-- Set stock if provided
	}

	if in.Price != nil {
		p.Price = *in.Price
	}

	// Only add variants if explicitly provided
	if len(in.Variants) > 0 {
		for _, v := range in.Variants {
			var opts []models.Option
			for _, o := range v.Options {
				opts = append(opts, models.Option{Name: o.Name, Value: o.Value})
			}
			p.Variants = append(p.Variants, models.Variant{
				Options: opts,
				Price:   v.Price,
				Stock:   v.Stock,
				Image:   v.Image,
			})
		}
	}

	_, err = services.CreateProductService(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creation failed: " + err.Error()})
		return
	}

	// Fetch the created product and return the correct API response
	createdProduct, _ := services.GetProductByIDService(p.ID.Hex())
	c.JSON(http.StatusCreated, services.ProductToAPIResponse(createdProduct))
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

	apiProducts := make([]map[string]interface{}, 0, len(products))
	for i := range products {
		apiProducts = append(apiProducts, services.ProductToAPIResponse(&products[i]))
	}
	c.JSON(http.StatusOK, apiProducts)
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
	c.JSON(http.StatusOK, services.ProductToAPIResponse(p))
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

	// Custom validation for price/variants on update
	if variants, ok := upd["variants"].([]interface{}); ok && len(variants) == 0 {
		if price, hasPrice := upd["price"]; hasPrice {
			if priceFloat, isFloat := price.(float64); !isFloat || priceFloat < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Product price must be a non-negative number if provided and no variants exist"})
				return
			}
		}
		// Do not require price/stock to be present for PATCH; only validate if present
	}

	// Validate variants if present
	if rawVariants, ok := upd["variants"].([]interface{}); ok && len(rawVariants) > 0 {
		for _, rv := range rawVariants {
			if vMap, isMap := rv.(map[string]interface{}); isMap {
				opts, hasOpts := vMap["options"].([]interface{})
				if !hasOpts || len(opts) == 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Each variant must have at least one option (name:value)"})
					return
				}
				for _, o := range opts {
					if oMap, ok := o.(map[string]interface{}); ok {
						name, _ := oMap["name"].(string)
						value, _ := oMap["value"].(string)
						if strings.TrimSpace(name) == "" || strings.TrimSpace(value) == "" {
							c.JSON(http.StatusBadRequest, gin.H{"error": "Variant option name and value cannot be empty"})
							return
						}
					}
				}
			}
		}
	}

	_, err = services.UpdateProductService(c.Param("productId"), upd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	// Fetch the updated product and return the correct API response
	updatedProduct, _ := services.GetProductByIDService(c.Param("productId"))
	c.JSON(http.StatusOK, services.ProductToAPIResponse(updatedProduct))
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
