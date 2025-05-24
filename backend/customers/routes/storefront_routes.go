package routes

import (
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

// StorefrontRoutes sets up public storefront endpoints.
func StorefrontRoutes(r *gin.Engine) {
	shops := r.Group("/shops")
	{
		// GET /shops/:shopid               → shop details
		shops.GET("/:shopid", controllers.GetStorefront)

		// Nested products under that shop...
		products := shops.Group("/:shopid/products")
		{
			// GET /shops/:shopid/products
			products.GET("", controllers.GetProductsByShop)

			// GET /shops/:shopid/products/:productid
			products.GET("/:productid", controllers.GetProductDetail)
		}
	}
}
