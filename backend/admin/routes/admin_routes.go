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
		productGroup.GET("/:id", middlewares.AdminAuthMiddleware(), controllers.GetProduct)
		productGroup.POST("/create", middlewares.AdminAuthMiddleware(), controllers.CreateProduct)
        productGroup.PATCH("/update/:id", middlewares.AdminAuthMiddleware(), controllers.UpdateProduct)
		productGroup.DELETE("/delete/:id", middlewares.AdminAuthMiddleware(), controllers.DeleteProduct)
		productGroup.GET("/count", middlewares.AdminAuthMiddleware(), controllers.GetProductCount)
		productGroup.GET("/count-by-category", middlewares.AdminAuthMiddleware(), controllers.GetProductsByCategoryCount)


	}

	
	//  Customers sub-group (Example)
	 customerGroup := adminGroup.Group("/customers")
	 {
	 	customerGroup.GET("/", middlewares.AdminAuthMiddleware(), controllers.GetCustomers)
		customerGroup.GET("/:id",  middlewares.AdminAuthMiddleware(), controllers.GetCustomer)
		customerGroup.POST("/create", middlewares.AdminAuthMiddleware(),  controllers.CreateCustomer)
		customerGroup.PATCH("/update/:id",  middlewares.AdminAuthMiddleware(),  controllers.UpdateCustomer)
		customerGroup.DELETE("/delete/:id",  middlewares.AdminAuthMiddleware(), controllers.DeleteCustomer)
	 	
	 }

	// Shops sub-group 
	 shopGroup := adminGroup.Group("/stores")
	 {
	 	shopGroup.GET("/", middlewares.AdminAuthMiddleware(), controllers.GetShops)
		shopGroup.GET("/:id", middlewares.AdminAuthMiddleware(), controllers.GetShop)
	 	shopGroup.POST("/create",middlewares.AdminAuthMiddleware(),  controllers.CreateShop)
		shopGroup.PATCH("/update/:id", middlewares.AdminAuthMiddleware(), controllers.UpdateShop)
		shopGroup.DELETE("/delete/:id", middlewares.AdminAuthMiddleware(), controllers.DeleteShop)

	 }

	// Order sub-group 
	 orderGroup := adminGroup.Group("/orders")
	 {
	 	orderGroup.GET("/",middlewares.AdminAuthMiddleware(),  controllers.GetOrders)
		orderGroup.GET("/:id", middlewares.AdminAuthMiddleware(), controllers.GetOrder)
	 	orderGroup.POST("/create", middlewares.AdminAuthMiddleware(),  controllers.CreateOrder)
		orderGroup.PATCH("/update/:id", middlewares.AdminAuthMiddleware(), controllers.UpdateOrder)
		orderGroup.DELETE("/delete/:id", middlewares.AdminAuthMiddleware(), controllers.DeleteOrder)
	 }

    // Sellers  sub-group (Example)
	 sellerGroup := adminGroup.Group("/sellers")
	 {
	 	sellerGroup.GET("/",middlewares.AdminAuthMiddleware(),  controllers.GetSellers)
	 	sellerGroup.POST("/create-seller", middlewares.AdminAuthMiddleware(),  controllers.CreateSeller)
		sellerGroup.PATCH("/update-shop", middlewares.AdminAuthMiddleware(), controllers.UpdateSeller)
	 }
}
