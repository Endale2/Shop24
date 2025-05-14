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
		productGroup.POST("/create-product", middlewares.AdminAuthMiddleware(), controllers.CreateProduct)
		
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

	// Shops sub-group (Example)
	 shopGroup := adminGroup.Group("/stores")
	 {
	 	shopGroup.GET("/", middlewares.AdminAuthMiddleware(), controllers.GetShops)
	 	shopGroup.POST("/create-store",middlewares.AdminAuthMiddleware(),  controllers.CreateShop)
		shopGroup.PATCH("/update-shop", middlewares.AdminAuthMiddleware(), controllers.UpdateShop)
	 }

	// Order sub-group (Example)
	 orderGroup := adminGroup.Group("/orders")
	 {
	 	orderGroup.GET("/",middlewares.AdminAuthMiddleware(),  controllers.GetOrders)
	 	orderGroup.POST("/create-order", middlewares.AdminAuthMiddleware(),  controllers.CreateOrder)
		orderGroup.PATCH("/update-order", middlewares.AdminAuthMiddleware(), controllers.UpdateOrder)
	 }

    // Sellers  sub-group (Example)
	 sellerGroup := adminGroup.Group("/sellers")
	 {
	 	sellerGroup.GET("/",middlewares.AdminAuthMiddleware(),  controllers.GetSellers)
	 	sellerGroup.POST("/create-seller", middlewares.AdminAuthMiddleware(),  controllers.CreateSeller)
		sellerGroup.PATCH("/update-shop", middlewares.AdminAuthMiddleware(), controllers.UpdateSeller)
	 }
}
