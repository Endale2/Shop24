package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/auth/controllers"
)

// RegisterAuthRoutes registers all authentication-related routes.
func RegisterAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/admin/login", controllers.AdminLogin)
		authGroup.POST("/admin/register", controllers.AdminRegister)
		// You can add routes for sellers and customers as needed.
	}
}
