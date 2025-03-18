package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/controllers"
)

// RegisterAuthRoutes registers all authentication-related routes.
func RegisterAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	
	{   // Admin authentication routes.
		authGroup.POST("/admin/login", controllers.AdminLogin)
		authGroup.POST("/admin/register", controllers.AdminRegister)
		// Seller authentication routes.
	    authGroup.POST("/seller/register", controllers.SellerRegister)
	    authGroup.POST("/seller/login", controllers.SellerLogin)
	}
}
