package routes

import (
	"github.com/Endale2/DRPS/customers/controllers"
	"github.com/gin-gonic/gin"
)

// StorefrontRoutes sets up public routes for storefront operations.
func StorefrontRoutes(r *gin.RouterGroup) {
	// Group storefront routes under /storefront
	storefront := r.Group("/storefront")
	{
		// GET /customers/storefront/:id
		storefront.GET("/:id", controllers.GetStorefront)
	}
}
