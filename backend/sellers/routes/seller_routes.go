package routes

import (
	"github.com/Endale2/DRPS/sellers/controllers"
	"github.com/Endale2/DRPS/auth/middlewares"

	"github.com/gin-gonic/gin"
)


func  SellerRoute(r  *gin.Engine){
   sellerGroup:=r.Group("/seller")

   //product  group

   productGroup:=sellerGroup.Group("/products")
   {
     productGroup.GET("/")
   }

}

