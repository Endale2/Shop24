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
		authGroup.POST("/admin/create", controllers.AdminRegister)
		authGroup.POST("/admin/oauth",    controllers.AdminOAuth)
        authGroup.POST("/admin/logout", controllers.AdminLogout)
		authGroup.POST("/admin/refresh", controllers.AdminRefresh)
		authGroup.GET("/admin/me", controllers.GetAuthAdminMe)

		// Seller authentication routes
		authGroup.POST("/seller/register", controllers.SellerRegister)
		authGroup.POST("/seller/login", controllers.SellerLogin)
		authGroup.GET("/seller/oauth/google",         controllers.SellerOAuthRedirect)
        authGroup.GET("/seller/oauth/google/callback", controllers.SellerOAuthCallback)
		authGroup.POST("/seller/refresh",  controllers.SellerRefresh)
		authGroup.POST("/seller/logout", controllers.SellerLogout)
		authGroup.GET("/seller/me", controllers.GetCurrentSeller)

		// Customer authentication routes
		authGroup.POST("/customer/register", controllers.CustomerRegister)
		authGroup.POST("/customer/login", controllers.CustomerLogin)
		authGroup.POST("/customer/refresh",  controllers.CustomerRefresh)
		authGroup.POST("/customer/logout", controllers.CustomerLogout)
		authGroup.GET("/customer/me", controllers.CustomerMe)
	}
}
