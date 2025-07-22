package routes

import (
	"github.com/Endale2/DRPS/admin/controllers"
	"github.com/Endale2/DRPS/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	adminGroup := r.Group("/admin")

	// Products sub-group
	productGroup := adminGroup.Group("/products", middlewares.AuthMiddleware())
	{
		productGroup.GET("/", controllers.GetProducts)
		productGroup.GET("/:id", controllers.GetProduct)
		productGroup.POST("/create", controllers.CreateProduct)
		productGroup.PATCH("/update/:id", controllers.UpdateProduct)
		productGroup.DELETE("/delete/:id", controllers.DeleteProduct)
		productGroup.GET("/count", controllers.GetProductCount)
	}

	//  Customers sub-group
	customerGroup := adminGroup.Group("/customers", middlewares.AuthMiddleware())
	{
		customerGroup.GET("/", controllers.GetCustomers)
		customerGroup.GET("/:id", controllers.GetCustomer)
		customerGroup.POST("/create", controllers.CreateCustomer)
		customerGroup.PATCH("/update/:id", controllers.UpdateCustomer)
		customerGroup.DELETE("/delete/:id", controllers.DeleteCustomer)

	}

	// Shops sub-group
	shopGroup := adminGroup.Group("/stores", middlewares.AuthMiddleware())
	{
		shopGroup.GET("/", controllers.GetShops)
		shopGroup.GET("/:id", controllers.GetShop)
		shopGroup.POST("/create", controllers.CreateShop)
		shopGroup.PATCH("/update/:id", controllers.UpdateShop)
		shopGroup.DELETE("/delete/:id", controllers.DeleteShop)

	}

	// Sellers  sub-group
	sellerGroup := adminGroup.Group("/sellers", middlewares.AuthMiddleware())
	{
		sellerGroup.GET("/", controllers.GetSellers)
		sellerGroup.GET("/:id", controllers.GetSeller)
		sellerGroup.POST("/create", controllers.CreateSeller)
		sellerGroup.PATCH("/update/:id", controllers.UpdateSeller)
		sellerGroup.DELETE("/delete/:id", controllers.DeleteSeller)
	}

	// Order sub-group
	orderGroup := adminGroup.Group("/orders", middlewares.AuthMiddleware())
	{
		orderGroup.GET("/", controllers.GetOrders)
		orderGroup.GET("/:id", controllers.GetOrder)
		orderGroup.POST("/create", controllers.CreateOrder)
		orderGroup.PATCH("/update/:id", controllers.UpdateOrder)
		orderGroup.DELETE("/delete/:id", controllers.DeleteOrder)
		orderGroup.PATCH("/update-status/:id", controllers.UpdateOrderStatus)
	}

	// Analytics sub-group
	analyticsGroup := adminGroup.Group("/analytics", middlewares.AuthMiddleware())
	{
		analyticsGroup.GET("/total-sales", controllers.GetTotalSales)
		analyticsGroup.GET("/total-revenue", controllers.GetTotalRevenue)
		analyticsGroup.GET("/order-count", controllers.GetOrderCount)
		analyticsGroup.GET("/top-products", controllers.GetTopProducts)
		analyticsGroup.GET("/top-customers", controllers.GetTopCustomers)
	}

	// Staff sub-group
	staffGroup := adminGroup.Group("/staff", middlewares.AuthMiddleware())
	{
		staffGroup.GET("/", controllers.GetAdmins)
		staffGroup.GET("/:id", controllers.GetAdmin)
		staffGroup.POST("/create", controllers.CreateAdmin)
		staffGroup.PATCH("/update/:id", controllers.UpdateAdmin)
		staffGroup.DELETE("/delete/:id", controllers.DeleteAdmin)
	}
}
