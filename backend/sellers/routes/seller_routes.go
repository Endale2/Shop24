package routes

import (
	"github.com/Endale2/DRPS/auth/middlewares"
	"github.com/Endale2/DRPS/sellers/controllers"
	"github.com/gin-gonic/gin"
)

func SellerRoute(r *gin.Engine) {
	// 1) Root seller group, with auth
	sellerGroup := r.Group("/seller")
	sellerGroup.Use(middlewares.AuthMiddleware())

	// 2) /seller/shops  — list & create
	sellerGroup.POST("/shops", controllers.CreateShop)
	sellerGroup.GET( "/shops", controllers.GetShops)

	// 3) /seller/shops/:shopId
	shopGroup := sellerGroup.Group("/shops/:shopId")
	{
		// shop CRUD
		shopGroup.GET(   "",               controllers.GetShop)
		shopGroup.PATCH( "",               controllers.UpdateShop)
		shopGroup.DELETE("",               controllers.DeleteShop)

		// nested products under that shop
		prodGroup := shopGroup.Group("/products")
		{
			prodGroup.POST(   "",                 controllers.CreateProduct)
			prodGroup.GET(    "",                 controllers.GetProducts)
			prodGroup.GET(    "/:productId",      controllers.GetProduct)
			prodGroup.PATCH(  "/:productId",      controllers.UpdateProduct)
			prodGroup.DELETE("/:productId",      controllers.DeleteProduct)
		}
	}
}
