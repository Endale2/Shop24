package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/controllers"
)

// AuthRoutes registers all authentication-related routes.
func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		// Admin authentication routes
		authGroup.POST("/admin/login", controllers.AdminLogin)
		authGroup.POST("/admin/register", controllers.AdminRegister)
		// Refresh access token using refresh token
		authGroup.POST("/admin/refresh", controllers.RefreshToken)
		authGroup.GET("/admin/me", controllers.GetAuthAdminMe)

		// Seller authentication routes
		authGroup.POST("/seller/register", controllers.SellerRegister)
		authGroup.POST("/seller/login", controllers.SellerLogin)

		// Customer authentication routes
		authGroup.POST("/customer/register", controllers.CustomerRegister)
		authGroup.POST("/customer/login", controllers.CustomerLogin)
	}
}
