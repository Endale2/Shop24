package routes

import (
	"github.com/Endale2/DRPS/auth/middlewares"
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

func StorefrontRoutes(r *gin.Engine) {
	// All storefront endpoints begin with /shops/:shopSlug
	shops := r.Group("/shops")
	{
		// GET /shops/:shopSlug → view shop details
		shops.GET("/:shopSlug", controllers.GetStorefront)

		// Collection endpoints nested under /shops/:shopSlug/collections
		collections := shops.Group("/:shopSlug/collections")
		{
			// list all collections for this shop
			collections.GET("", controllers.ListCollections)
			// get one collection + its product summaries
			collections.GET("/:collectionHandle", controllers.GetCollection)
		}

		// Product endpoints nested under /shops/:shopSlug/products
		products := shops.Group("/:shopSlug/products")
		{
			products.GET("", controllers.GetProductsByShop)
			products.GET("/:productSlug", controllers.GetProductDetail)
		}

		// Protected endpoints: all under /shops/:shopSlug/ and require auth
		auth := shops.Group("/:shopSlug", middlewares.AuthMiddleware())
		{
			auth.GET("/cart", controllers.GetCart)
			auth.POST("/cart/items", controllers.AddToCart)
			auth.PUT("/cart/items", controllers.UpdateCartItem)
			auth.DELETE("/cart/items", controllers.RemoveCartItem)
			auth.POST("/cart/discount", controllers.ApplyCartDiscount)
			auth.POST("/cart/clear", controllers.ClearCart)
			auth.POST("/orders", controllers.PlaceOrder)
			auth.GET("/orders", controllers.ListShopOrders)
			auth.GET("/orders/:orderId", controllers.GetOrderDetail)
		}
	}
}
