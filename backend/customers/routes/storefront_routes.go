package routes

import (
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

// StorefrontRoutes sets up public routes for storefront operations.
func StorefrontRoutes(r *gin.Engine) {
	// Group storefront routes under /storefront
	storefront := r.Group("/shops")
	{
		// GET /customers/storefront/:id
		storefront.GET("/:id", controllers.GetStorefront)
	}
}
