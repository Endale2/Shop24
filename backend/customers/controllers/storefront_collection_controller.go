package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/services"
	"github.com/Endale2/DRPS/sellers/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CollectionProductSummary is the minimal info we expose for each product.
type CollectionProductSummary struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	MainImage string             `json:"main_image"`
}

// ListCollections handles GET /shops/:shopSlug/collections
func ListCollections(c *gin.Context) {
	slug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// ← here we call GetCollectionsByShop, not ListCollectionsByShopID
	cols, err := repositories.GetCollectionsByShop(shop.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list collections"})
		return
	}

	out := make([]gin.H, 0, len(cols))
	for _, coll := range cols {
		out = append(out, gin.H{
			"id":     coll.ID,
			"title":  coll.Title,
			"handle": coll.Handle,
			"image":  coll.Image,
		})
	}

	c.JSON(http.StatusOK, out)
}

// GetCollection handles GET /shops/:shopSlug/collections/:collId
func GetCollection(c *gin.Context) {
	slug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	collIDHex := c.Param("collId")
	collID, err := primitive.ObjectIDFromHex(collIDHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid collection ID"})
		return
	}

	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving collection"})
		return
	}
	if coll == nil || coll.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "collection not found"})
		return
	}

	// Build summaries of products
	summaries := []CollectionProductSummary{}
	for _, pid := range coll.ProductIDs {
		p, err := services.GetProductByIDService(pid.Hex())
		if err != nil || p == nil {
			continue
		}
		summaries = append(summaries, CollectionProductSummary{
			ID:        p.ID,
			Name:      p.Name,
			MainImage: p.MainImage,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          coll.ID,
		"title":       coll.Title,
		"description": coll.Description,
		"handle":      coll.Handle,
		"image":       coll.Image,
		"created_at":  coll.CreatedAt,
		"filters":     coll.Filters,
		"products":    summaries,
	})
}
