// customers/controllers/storefront_product_controller.go
package controllers

import (
	"net/http"

	"strconv"

	"strings"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProductsByShop returns all products for a shop identified by its slug.
// GET /shops/:shopSlug/products
func GetProductsByShop(c *gin.Context) {
	shopSlug := c.Param("shopSlug")

	// 1) Verify shop exists
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// 2) Parse pagination params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	skip := (page - 1) * limit

	// 3) Fetch products by shop slug (service filters by shop_id internally)
	products, err := services.GetProductsByShopSlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	total := len(products)
	// Paginate in-memory (for now; ideally, do this in the DB query)
	end := skip + limit
	if end > total {
		end = total
	}
	pagedProducts := products
	if skip < total {
		pagedProducts = products[skip:end]
	} else {
		pagedProducts = []models.Product{}
	}

	// Patch: Ensure variant options is never null
	apiProducts := make([]map[string]interface{}, 0, len(pagedProducts))
	for i := range pagedProducts {
		for j := range pagedProducts[i].Variants {
			if pagedProducts[i].Variants[j].Options == nil {
				pagedProducts[i].Variants[j].Options = []models.Option{}
			}
		}
		apiProducts = append(apiProducts, services.ProductToAPIResponseWithDiscounts(&pagedProducts[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"products": apiProducts,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

// GetProductDetail returns one product by its slug, verifying it belongs to this shop.
// GET /shops/:shopSlug/products/:productSlug
func GetProductDetail(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	productSlug := c.Param("productSlug")

	// 1) Verify shop exists
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// 2) Fetch product by slug
	product, err := services.GetProductBySlugService(productSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve product"})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	// 3) Ensure product.ShopID matches shop.ID
	if product.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "product does not belong to this shop"})
		return
	}

	c.JSON(http.StatusOK, services.ProductToAPIResponseWithDiscounts(product))
}

// GET /shops/:shopSlug/products/id/:productId
func GetProductDetailByID(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	productId := c.Param("productId")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil || shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}
	pid, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}
	product, err := services.GetProductByIDService(pid.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil || product.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// SearchProductsByShop returns paginated products matching a search query for a shop
// GET /shops/:shopSlug/products/search?q=...&page=1&limit=20
func SearchProductsByShop(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	q := strings.ToLower(c.DefaultQuery("q", ""))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	skip := (page - 1) * limit

	// 1) Verify shop exists
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// 2) Fetch all products for the shop
	products, err := services.GetProductsByShopSlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	// 3) Filter by search query (name or description, case-insensitive)
	var filtered []models.Product
	if q != "" {
		for _, p := range products {
			if strings.Contains(strings.ToLower(p.Name), q) || strings.Contains(strings.ToLower(p.Description), q) {
				filtered = append(filtered, p)
			}
		}
	} else {
		filtered = products
	}

	total := len(filtered)
	end := skip + limit
	if end > total {
		end = total
	}
	pagedProducts := filtered
	if skip < total {
		pagedProducts = filtered[skip:end]
	} else {
		pagedProducts = []models.Product{}
	}

	apiProducts := make([]map[string]interface{}, 0, len(pagedProducts))
	for i := range pagedProducts {
		for j := range pagedProducts[i].Variants {
			if pagedProducts[i].Variants[j].Options == nil {
				pagedProducts[i].Variants[j].Options = []models.Option{}
			}
		}
		apiProducts = append(apiProducts, services.ProductToAPIResponseWithDiscounts(&pagedProducts[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"products": apiProducts,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}
