// customers/controllers/order_controller.go
package controllers

import (
	"net/http"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlaceOrder handles POST /shops/:shopSlug/orders
func PlaceOrder(c *gin.Context) {
	// 1) Lookup the shop by its slug
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// 2) Bind the incoming order JSON into our Order model
	var in models.Order
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3) Stamp the shop ID
	in.ShopID = shop.ID

	// 4) If you have customer auth, pull the customer ID from context:
	if cidVal, exists := c.Get("customer_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				in.CustomerID = cid
			}
		}
	}

	// 5) Create the order
	created, err := services.CreateOrderService(&in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// ListShopOrders handles GET /shops/:shopSlug/orders
func ListShopOrders(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	// List by shop ID
	list, err := services.ListOrdersByShopService(shop.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Filter to only the current customer, if logged in
	if cidVal, exists := c.Get("customer_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				filtered := make([]models.Order, 0, len(list))
				for _, o := range list {
					if o.CustomerID == cid {
						filtered = append(filtered, o)
					}
				}
				list = filtered
			}
		}
	}

	c.JSON(http.StatusOK, list)
}

// GetOrderDetail handles GET /shops/:shopSlug/orders/:orderId
func GetOrderDetail(c *gin.Context) {
	shopSlug := c.Param("shopSlug")
	shop, err := services.GetShopBySlugService(shopSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup shop"})
		return
	}
	if shop == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "shop not found"})
		return
	}

	orderID := c.Param("orderId")
	order, err := services.GetOrderByIDService(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order == nil || order.ShopID != shop.ID {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	// Only the placing customer may view it
	if cidVal, exists := c.Get("customer_id"); exists {
		if cidHex, ok := cidVal.(string); ok {
			if cid, err := primitive.ObjectIDFromHex(cidHex); err == nil {
				if order.CustomerID != cid {
					c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to view this order"})
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, order)
}
