package routes

import (
	"github.com/Endale2/DRPS/auth/middlewares"
	"github.com/Endale2/DRPS/sellers/controllers"
	"github.com/gin-gonic/gin"
)

func SellerRoute(r *gin.Engine) {
	// 1) Root seller group, with auth
	sellerGroup := r.Group("/seller")
	sellerGroup.Use(middlewares.AuthMiddleware())

	// 2) /seller/shops  — list & create
	sellerGroup.POST("/shops", controllers.CreateShop)
	sellerGroup.GET("/shops", controllers.GetShops)

	// 3) /seller/categories — get all categories
	sellerGroup.GET("/categories", controllers.GetShopCategories)

	// ═════════  Theme Management Routes ═════════
	// Base theme operations
	themeGroup := sellerGroup.Group("/themes")
	{
		themeGroup.POST("", controllers.CreateTheme)                     // Create custom theme
		themeGroup.GET("/marketplace", controllers.GetPublicThemes)      // List marketplace themes
		themeGroup.GET("/my-themes", controllers.GetMyThemes)            // List seller's custom themes
		themeGroup.GET("/search", controllers.SearchThemes)              // Search themes
		themeGroup.GET("/:themeId", controllers.GetTheme)                // Get specific theme
		themeGroup.PATCH("/:themeId", controllers.UpdateTheme)           // Update theme
		themeGroup.DELETE("/:themeId", controllers.DeleteTheme)          // Delete theme
		themeGroup.GET("/:themeId/presets", controllers.GetThemePresets) // Get theme presets
		themeGroup.POST("/seed", controllers.SeedDefaultThemes)          // Seed default themes (dev/admin)
	}

	// 3) /seller/shops/:shopId
	shopGroup := sellerGroup.Group("/shops/:shopId")
	{
		// shop CRUD
		shopGroup.GET("", controllers.GetShop)
		shopGroup.PATCH("", controllers.UpdateShop)
		shopGroup.DELETE("", controllers.DeleteShop)

		// nested products under that shop
		prodGroup := shopGroup.Group("/products")
		{
			prodGroup.POST("", controllers.CreateProduct)
			prodGroup.GET("", controllers.GetProducts)
			prodGroup.GET("/:productId", controllers.GetProduct)
			prodGroup.PATCH("/:productId", controllers.UpdateProduct)
			prodGroup.DELETE("/:productId", controllers.DeleteProduct)
		}
		// ─────  nested customers  routes ─────
		custGroup := shopGroup.Group("/customers")
		{
			custGroup.POST("/link", controllers.LinkCustomer)
			custGroup.GET("", controllers.GetLinkedCustomers)
			custGroup.GET("/stats", controllers.GetCustomerStats)
			custGroup.GET("/:customerId", controllers.GetCustomerDetail)
			custGroup.DELETE("/link/:linkId", controllers.UnlinkCustomer)
			custGroup.GET("/:customerId/history", controllers.GetCustomerOrderHistory)
		}

		// ─────  nested customer segments  routes ─────
		segmentGroup := shopGroup.Group("/customer-segments")
		{
			segmentGroup.POST("", controllers.CreateCustomerSegment)
			segmentGroup.GET("", controllers.GetCustomerSegments)
			segmentGroup.PATCH("/:segmentId", controllers.UpdateCustomerSegment)
			segmentGroup.DELETE("/:segmentId", controllers.DeleteCustomerSegment)

			// Add/remove customers from segments
			segmentGroup.POST("/:segmentId/customers", controllers.AddCustomerToSegment)
			segmentGroup.DELETE("/:segmentId/customers/:customerId", controllers.RemoveCustomerFromSegment)
		}

		// ─────  nested collections ─────
		collGroup := shopGroup.Group("/collections")
		{
			// Create a new collection
			collGroup.POST("", controllers.CreateCollection)

			// List all collections for this shop
			collGroup.GET("", controllers.GetCollections)

			// CRUD on a single collection
			collGroup.GET("/:collId", controllers.GetCollection)
			collGroup.PATCH("/:collId", controllers.UpdateCollection)
			collGroup.PUT("/:collId", controllers.UpdateCollection)
			collGroup.DELETE("/:collId", controllers.DeleteCollection)

			// Add/remove products in a given collection
			collGroup.POST("/:collId/products", controllers.AddProductsToCollection)
			collGroup.DELETE("/:collId/products", controllers.RemoveProductsFromCollection)
		}

		// ─────  nested Discount  routes ─────
		discGroup := shopGroup.Group("/discounts")
		{
			discGroup.POST("", controllers.CreateDiscount)
			discGroup.GET("", controllers.ListDiscounts)
			discGroup.GET("/:id", controllers.GetDiscount)
			discGroup.PATCH("/:id", controllers.UpdateDiscount)
			discGroup.DELETE("/:id", controllers.DeleteDiscount)

			// Enhanced discount functionality
			discGroup.POST("/:id/customers", controllers.AddCustomersToDiscount)
			discGroup.POST("/:id/segments", controllers.AddSegmentsToDiscount)
			discGroup.POST("/:id/mixed-eligibility", controllers.AddMixedEligibility)
			discGroup.POST("/:id/clear-eligibility", controllers.ClearAllowedCustomersAndSegments)
			discGroup.GET("/:id/eligible-customers", controllers.GetEligibleCustomers)
			discGroup.GET("/:id/usage", controllers.GetDiscountUsageStats)
			discGroup.POST("/:id/validate", controllers.ValidateDiscountForCustomer)

			// Paginated products under discount
			discGroup.GET("/:id/products", controllers.GetDiscountProductsPaginated)
		}

		//orders
		orders := shopGroup.Group("/orders")
		{
			orders.POST("", controllers.CreateOrder)
			orders.GET("", controllers.ListOrders)
			orders.GET("/dashboard", controllers.GetOrdersForDashboard)
			orders.GET("/stats", controllers.GetOrderStats)
			orders.GET("/:orderId", controllers.GetOrder)
			orders.GET("/:orderId/details", controllers.GetOrderWithCustomerDetails)
			orders.PATCH("/:orderId", controllers.UpdateOrder)
			orders.DELETE("/:orderId", controllers.DeleteOrder)
		}

		// ─────  Analytics endpoints ─────
		analyticsGroup := shopGroup.Group("/analytics")
		{
			analyticsGroup.GET("/summary", controllers.GetShopAnalyticsSummary)
			analyticsGroup.GET("/revenue-over-time", controllers.GetShopRevenueOverTime)
			analyticsGroup.GET("/orders-over-time", controllers.GetShopOrdersOverTime)
			analyticsGroup.GET("/top-products", controllers.GetShopTopProducts)
			analyticsGroup.GET("/top-customers", controllers.GetShopTopCustomers)
			analyticsGroup.GET("/inventory-status", controllers.GetShopInventoryStatus)
			analyticsGroup.GET("/discount-performance", controllers.GetShopDiscountPerformance)
			analyticsGroup.GET("/customers-over-time", controllers.GetShopCustomersOverTime)
			analyticsGroup.GET("/category-sales", controllers.GetShopCategorySales)
			analyticsGroup.GET("/recent-orders", controllers.GetShopRecentOrders)
			analyticsGroup.GET("/dashboard", controllers.GetShopDashboardAnalytics) // NEW ENDPOINT
		}

		// ═════════  Shop Theme & Customization Routes ═════════
		shopThemeGroup := shopGroup.Group("/themes")
		{
			shopThemeGroup.POST("", controllers.CreateShopTheme)                                       // Create shop theme config
			shopThemeGroup.GET("", controllers.GetShopThemes)                                          // List shop theme configs
			shopThemeGroup.GET("/active", controllers.GetShopActiveTheme)                              // Get active theme
			shopThemeGroup.PATCH("/:shopThemeId", controllers.UpdateShopTheme)                         // Update shop theme
			shopThemeGroup.POST("/:shopThemeId/publish", controllers.PublishShopTheme)                 // Publish/activate theme
			shopThemeGroup.DELETE("/:shopThemeId", controllers.DeleteShopTheme)                        // Delete shop theme config
			shopThemeGroup.POST("/:shopThemeId/presets/:presetId/apply", controllers.ApplyThemePreset) // Apply preset
		}

		// ─────  Customization API endpoints ─────
		customizationGroup := shopGroup.Group("/customization")
		{
			customizationGroup.GET("", controllers.GetCustomizationData)  // Get current customization
			customizationGroup.PATCH("", controllers.UpdateCustomization) // Update customization (full)

			// Granular theme updates (matching frontend API calls)
			customizationGroup.PATCH("/colors", controllers.SaveThemeColors)  // Save colors only
			customizationGroup.PATCH("/fonts", controllers.SaveTypography)    // Save fonts only
			customizationGroup.PATCH("/layout", controllers.SaveLayout)       // Save layout only
			customizationGroup.POST("/reset", controllers.ResetCustomization) // Reset to defaults

			// Enhanced styling endpoints
			customizationGroup.PATCH("/gradients", controllers.SaveGradients)   // Save gradients
			customizationGroup.PATCH("/shadows", controllers.SaveShadows)       // Save shadows
			customizationGroup.PATCH("/animations", controllers.SaveAnimations) // Save animations
			customizationGroup.PATCH("/components", controllers.SaveComponents) // Save component configs
		}

	}

}
