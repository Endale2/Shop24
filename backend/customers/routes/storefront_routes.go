// customers/routes/storefront_routes.go
package routes

import (
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

func StorefrontRoutes(r *gin.Engine) {
	// All storefront endpoints begin with /shops/:shopSlug
	shops := r.Group("/shops")
	{
		// GET /shops/:shopSlug → view shop details
		shops.GET("/:shopSlug", controllers.GetStorefront)

		// Product endpoints nested under /shops/:shopSlug/products
		products := shops.Group("/:shopSlug/products")
		{
			products.GET("", controllers.GetProductsByShop)
			products.GET("/:productSlug", controllers.GetProductDetail)
		}

		// Cart endpoints nested under /shops/:shopSlug/cart
		cart := shops.Group("/:shopSlug/cart")
		{
			cart.GET("", controllers.GetOrCreateCart)
			cart.POST("/items", controllers.AddItemToCart)
			cart.PATCH("/items/:itemID", controllers.UpdateCartItem)
			cart.DELETE("/items/:itemID", controllers.RemoveCartItem)
			cart.POST("/clear", controllers.ClearCart)
			cart.POST("/apply-coupon", controllers.ApplyCoupon)
			cart.POST("/checkout", controllers.CheckoutCart)
		}
	}
}
