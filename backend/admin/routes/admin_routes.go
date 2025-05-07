package routes

import (
	"github.com/Endale2/DRPS/admin/controllers"
	"github.com/Endale2/DRPS/auth/middlewares"

	"github.com/gin-gonic/gin"
)

// AdminRoutes groups all admin-related routes under /admin
func AdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("/admin") // General /admin group

	// Products sub-group
	productGroup := adminGroup.Group("/products")
	{
		productGroup.GET("/", middlewares.AdminAuthMiddleware(), controllers.GetProducts)
		productGroup.POST("/create-product", controllers.CreateProduct)
	}

	
	// // Customers sub-group (Example)
	// customerGroup := adminGroup.Group("/customers")
	// {
	// 	customerGroup.GET("/", controllers.GetCustomers)
	// 	customerGroup.POST("/create-customer", controllers.CreateCustomer)
	// }

	// // Shops sub-group (Example)
	// shopGroup := adminGroup.Group("/shops")
	// {
	// 	shopGroup.GET("/", controllers.GetShops)
	// 	shopGroup.POST("/create-shop", controllers.CreateShop)
	// }
}
