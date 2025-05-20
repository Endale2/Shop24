package routes

import (
	 "github.com/Endale2/DRPS/sellers/controllers"
	"github.com/Endale2/DRPS/auth/middlewares"

	"github.com/gin-gonic/gin"
)


func  SellerRoute(r  *gin.Engine){
   sellerGroup:=r.Group("/seller", middlewares.AuthMiddleware())

   shops := sellerGroup.Group("/shops")
{
  shops.POST(   "/",            controllers.CreateShop)
  shops.GET(    "/",            controllers.GetShops)
  shops.GET(    "/:id",         controllers.GetShop)
  shops.PATCH(  "/:id",         controllers.UpdateShop)
  shops.DELETE( "/:id",         controllers.DeleteShop)

  prods := shops.Group("/:shopId/products")
  {
    prods.POST(   "/",            controllers.CreateProduct)
    prods.GET(    "/",            controllers.GetProducts)
    prods.GET(    "/:id",         controllers.GetProduct)
    prods.PATCH(  "/:id",         controllers.UpdateProduct)
    prods.DELETE( "/:id",         controllers.DeleteProduct)
  }
}


}

