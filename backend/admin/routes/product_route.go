package routes

import (
	"github.com/Endale2/DRPS/admin/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine){
	productGroup:=r.Group("products")

	{
     productGroup.GET("/", controllers.GetProducts)
	}
}