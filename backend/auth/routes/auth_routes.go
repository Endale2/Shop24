package routes

import (
	"github.com/Endale2/DRPS/auth/controllers"
	"github.com/Endale2/DRPS/auth/middlewares"
	"github.com/gin-gonic/gin"
)

// AuthRoutes registers all authentication-related routes.
func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		// Admin authentication routes
		authGroup.POST("/admin/login", controllers.AdminLogin)
		authGroup.POST("/admin/create", controllers.AdminRegister)
		authGroup.POST("/admin/oauth", controllers.AdminOAuth)
		authGroup.POST("/admin/logout", controllers.AdminLogout)
		authGroup.POST("/admin/refresh", controllers.AdminRefresh)
		authGroup.GET("/admin/me", controllers.GetAuthAdminMe)

		// Seller authentication routes
		authGroup.GET("/seller/oauth/google", controllers.SellerOAuthRedirect)
		authGroup.GET("/seller/oauth/google/callback", controllers.SellerOAuthCallback)

		authGroup.POST("/seller/refresh", controllers.SellerRefresh)
		authGroup.POST("/seller/logout", controllers.SellerLogout)
		authGroup.GET("/seller/me", controllers.GetCurrentSeller)
		authGroup.PATCH("/seller/update-profile", middlewares.AuthMiddleware(), controllers.UpdateCurrentSeller)

		// Customer authentication routes
		authGroup.POST("/customer/request-otp", controllers.CustomerRequestOTP)
		authGroup.POST("/customer/verify-otp", controllers.CustomerVerifyOTP)
		authGroup.POST("/customer/refresh", controllers.CustomerRefresh)
		authGroup.POST("/customer/logout", controllers.CustomerLogout)
		authGroup.GET("/customer/me", controllers.CustomerMe)
		authGroup.PATCH("/customer/me", controllers.UpdateCustomerMe)
	}
}
