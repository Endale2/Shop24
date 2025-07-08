package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CollectionProductSummary is the richer product info we expose in a collection.
type CollectionProductSummary struct {
	ID           primitive.ObjectID `json:"id"`
	Slug         string             `json:"slug"`
	Name         string             `json:"name"`
	MainImage    string             `json:"main_image"`
	Description  string             `json:"description"`
	Price        float64            `json:"price"`
	VariantCount int                `json:"variant_count"`
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

// GetCollection handles GET /shops/:shopSlug/collections/:collectionHandle
func GetCollection(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	handle := c.Param("collectionHandle")
	coll, err := repositories.GetCollectionByHandle(shop.ID, handle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving collection"})
		return
	}
	if coll == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "collection not found"})
		return
	}

	// Build richer product summaries
	var products []CollectionProductSummary
	for _, pid := range coll.ProductIDs {
		p, _ := services.GetProductByIDService(pid.Hex())
		if p == nil {
			continue
		}

		products = append(products, CollectionProductSummary{
			ID:           p.ID,
			Slug:         p.Slug,
			Name:         p.Name,
			MainImage:    p.MainImage,
			Description:  p.Description,
			Price:        p.Price,
			VariantCount: len(p.Variants),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          coll.ID,
		"title":       coll.Title,
		"handle":      coll.Handle,
		"description": coll.Description,
		"image":       coll.Image,
		"created_at":  coll.CreatedAt,
		"filters":     coll.Filters,
		"products":    products,
	})
}
