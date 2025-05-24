// routes/storefront_routes.go

package routes

import (
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

func StorefrontRoutes(r *gin.Engine) {
	shops := r.Group("/shops")
	{
		shops.GET("/:shopid", controllers.GetStorefront) // shop detail
		// products nested under the same :shopid
		prods := shops.Group("/:shopid/products")
		{
			prods.GET("", controllers.GetProductsByShop)
			prods.GET("/:productid", controllers.GetProductDetail)
		}
	}
}
