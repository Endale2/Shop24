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
		// GET /shops/:shopSlug → view shop details (legacy)
		shops.GET("/:shopSlug", controllers.GetStorefront)
		
		// GET /shops/:shopSlug/storefront → dynamic theme-driven storefront configuration
		shops.GET("/:shopSlug/storefront", controllers.GetDynamicStorefront)
		
		// GET /shops/:shopSlug/theme → real-time theme updates
		shops.GET("/:shopSlug/theme", controllers.GetStorefrontTheme)

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
			products.GET("search", controllers.SearchProductsByShop)
			products.GET("/:productSlug", controllers.GetProductDetail)
			products.GET("id/:productId", controllers.GetProductDetailByID)
		}

		// Debug endpoints (no auth required for testing)
		shops.GET("/:shopSlug/debug/product/:productId", controllers.DebugProduct)
		shops.POST("/:shopSlug/debug/product/:productId/price", controllers.UpdateProductPrice)
		shops.POST("/:shopSlug/debug/fix-variant-ids", controllers.FixVariantIDs)
		shops.GET("/:shopSlug/test-discounts", controllers.TestDiscounts)

		// Protected endpoints: all under /shops/:shopSlug/ and require auth
		auth := shops.Group("/:shopSlug", middlewares.AuthMiddleware())
		{
			auth.GET("/cart", controllers.GetCart)
			auth.POST("/cart/items", controllers.AddToCart)
			auth.PUT("/cart/items", controllers.UpdateCartItem)
			auth.DELETE("/cart/items", controllers.RemoveCartItem)
			auth.POST("/cart/clear", controllers.ClearCart)
			auth.POST("/orders", controllers.PlaceOrder)
			auth.GET("/orders", controllers.ListShopOrders)
			auth.GET("/orders/:orderId", controllers.GetOrderDetail)
			auth.GET("/wishlist", controllers.GetWishlist)
			auth.POST("/wishlist", controllers.AddToWishlist)
			auth.DELETE("/wishlist/:productId", controllers.RemoveFromWishlist)
		}
	}
}
