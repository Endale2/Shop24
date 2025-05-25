// routes/storefront_routes.go

package routes

import (
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

func StorefrontRoutes(r *gin.Engine) {
	shops := r.Group("/shops")
{
    shops.GET("/:slug", controllers.GetStorefront)
    prods := shops.Group("/:slug/products")
    {
        prods.GET("", controllers.GetProductsByShop)
        prods.GET("/:productslug", controllers.GetProductDetail)
    }
}

}
