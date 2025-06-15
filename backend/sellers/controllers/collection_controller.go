package controllers

import (
	"net/http"

	
	sharedSvc "github.com/Endale2/DRPS/shared/services"
	"github.com/Endale2/DRPS/sellers/services"
	"github.com/Endale2/DRPS/sellers/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collSvc = services.NewCollectionService()

// CreateCollectionInput is the expected JSON body for creating a collection.
type CreateCollectionInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Handle      string `json:"handle" binding:"required"`
	Image       string `json:"image"` // optional
}

// CreateCollection handles POST /seller/shops/:shopId/collections
func CreateCollection(c *gin.Context) {
	// 1) Get sellerID from JWT context
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	// 2) Parse shopId from URL
	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// 3) Bind body
	var in CreateCollectionInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// 4) Verify that shop belongs to seller
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	// 5) Call service
	newColl, err := collSvc.CreateCollectionService(
		shopID,
		in.Title,
		in.Description,
		in.Handle,
		in.Image,
		
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create collection"})
		return
	}
	c.JSON(http.StatusCreated, newColl)
}

// GetCollections handles GET /seller/shops/:shopId/collections
func GetCollections(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}

	// Verify ownership
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	colls, err := collSvc.GetCollectionsByShopService(shopID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch collections"})
		return
	}
	c.JSON(http.StatusOK, colls)
}

// GetCollection handles GET /seller/shops/:shopId/collections/:collId
func GetCollection(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopHex := c.Param("shopId")
	shopID, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	collHex := c.Param("collId")
	collID, err := primitive.ObjectIDFromHex(collHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid collection ID"})
		return
	}

	coll, err := repositories.GetCollectionByID(collID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving collection"})
		return
	}
	if coll == nil || coll.ShopID != shopID {
		c.JSON(http.StatusNotFound, gin.H{"error": "collection not found"})
		return
	}

	// Build a list of product summaries (id, name, main_image)
	type ProductSummary struct {
		ID        primitive.ObjectID `json:"id"`
		Name      string             `json:"name"`
		MainImage string             `json:"main_image"`
	}
	var summaries []ProductSummary
	for _, pid := range coll.ProductIDs {
		p, err := sharedSvc.GetProductByIDService(pid.Hex())
		if err != nil {
			// skip on error
			continue
		}
		if p == nil {
			continue
		}
		summaries = append(summaries, ProductSummary{
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
		"created_at":   coll.CreatedAt,
		"filters":      coll.Filters,
		"products":    summaries,
	})
}


// UpdateCollectionInput is the JSON body for updating a collection.
type UpdateCollectionInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Handle      *string `json:"handle"`
	Image       *string `json:"image"`
}

// UpdateCollection handles PATCH /seller/shops/:shopId/collections/:collId
func UpdateCollection(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopHex := c.Param("shopId")
	_, err := primitive.ObjectIDFromHex(shopHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
		return
	}
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	collHex := c.Param("collId")
	collID, err := primitive.ObjectIDFromHex(collHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid collection ID"})
		return
	}

	var in UpdateCollectionInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	updates := bson.M{}
	if in.Title != nil {
		updates["title"] = *in.Title
	}
	if in.Description != nil {
		updates["description"] = *in.Description
	}
	if in.Handle != nil {
		updates["handle"] = *in.Handle
	}
	if in.Image != nil {
		updates["image"] = *in.Image
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	err = collSvc.UpdateCollectionService(collID, sellerID, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update collection"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "collection updated"})
}

// DeleteCollection handles DELETE /seller/shops/:shopId/collections/:collId
func DeleteCollection(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	collHex := c.Param("collId")
	collID, err := primitive.ObjectIDFromHex(collHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid collection ID"})
		return
	}

	err = collSvc.DeleteCollectionService(collID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete collection"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "collection deleted"})
}

// AddProductInput is the JSON body for adding a product to a collection.
type AddProductInput struct {
	ProductID string `json:"productId" binding:"required"`
}

// AddProductToCollection handles POST /seller/shops/:shopId/collections/:collId/products
func AddProductToCollection(c *gin.Context) {
    userHex, _ := c.Get("user_id")
    sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

    // 1) Verify shop ownership
    shopHex := c.Param("shopId")
    shopID, err := primitive.ObjectIDFromHex(shopHex)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid shop ID"})
        return
    }
    shop, err := sharedSvc.GetShopByIDService(shopHex)
    if err != nil || shop.OwnerID != sellerID {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
        return
    }

    // 2) Parse collection ID
    collHex := c.Param("collId")
    collID, err := primitive.ObjectIDFromHex(collHex)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid collection ID"})
        return
    }
    coll, err := repositories.GetCollectionByID(collID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving collection"})
        return
    }
    if coll == nil || coll.ShopID != shopID {
        c.JSON(http.StatusNotFound, gin.H{"error": "collection not found"})
        return
    }

    // 3) Bind request payload
    var in AddProductInput
    if err := c.ShouldBindJSON(&in); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    prodID, err := primitive.ObjectIDFromHex(in.ProductID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
        return
    }

    // 4) Verify product exists and belongs to this shop
    prod, err := sharedSvc.GetProductByIDService(prodID.Hex())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching product"})
        return
    }
    if prod == nil || prod.ShopID != shopID {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product does not belong to this shop"})
        return
    }

    // 5) Prevent duplicates: check if already in coll.ProductIDs
    for _, existing := range coll.ProductIDs {
        if existing == prodID {
            c.JSON(http.StatusBadRequest, gin.H{"error": "product already in collection"})
            return
        }
    }

    // 6) Add product to collection
    if err := collSvc.AddProductToCollectionService(collID, prodID, sellerID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add product to collection"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "product added to collection"})
}

// RemoveProductFromCollection handles DELETE /seller/shops/:shopId/collections/:collId/products/:productId
func RemoveProductFromCollection(c *gin.Context) {
	userHex, _ := c.Get("user_id")
	sellerID, _ := primitive.ObjectIDFromHex(userHex.(string))

	shopHex := c.Param("shopId")
	shop, err := sharedSvc.GetShopByIDService(shopHex)
	if err != nil || shop.OwnerID != sellerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	collHex := c.Param("collId")
	collID, err := primitive.ObjectIDFromHex(collHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid collection ID"})
		return
	}

	prodHex := c.Param("productId")
	prodID, err := primitive.ObjectIDFromHex(prodHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	err = collSvc.RemoveProductFromCollectionService(collID, prodID, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove product from collection"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed from collection"})
}
