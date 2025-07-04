// customers/controllers/storefront_product_controller.go
package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
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

	// 2) Fetch products by shop slug (service filters by shop_id internally)
	products, err := services.GetProductsByShopSlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}
	// Patch: Ensure variant options is never null
	apiProducts := make([]map[string]interface{}, 0, len(products))
	for i := range products {
		for j := range products[i].Variants {
			if products[i].Variants[j].Options == nil {
				products[i].Variants[j].Options = []models.Option{}
			}
		}
		apiProducts = append(apiProducts, services.ProductToAPIResponse(&products[i]))
	}
	// It's valid for a shop to have zero products; return empty array
	c.JSON(http.StatusOK, apiProducts)
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

	c.JSON(http.StatusOK, services.ProductToAPIResponse(product))
}
