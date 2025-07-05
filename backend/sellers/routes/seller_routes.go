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
			custGroup.GET("/:customerId", controllers.GetCustomerDetail)
			custGroup.DELETE("/:linkId", controllers.UnlinkCustomer)
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
			collGroup.POST("/:collId/products", controllers.AddProductToCollection)
			collGroup.DELETE("/:collId/products/:productId", controllers.RemoveProductFromCollection)
		}

		// ─────  nested Discount  routes ─────
		discGroup := shopGroup.Group("/discounts")
		{
			discGroup.POST("", controllers.CreateDiscount)
			discGroup.GET("", controllers.ListDiscounts)
			discGroup.GET("/:id", controllers.GetDiscount)
			discGroup.PATCH("/:id", controllers.UpdateDiscount)
			discGroup.DELETE("/:id", controllers.DeleteDiscount)
		}

		//orders
		orders := shopGroup.Group("/orders")
		{
			orders.POST("", controllers.CreateOrder)
			orders.GET("", controllers.ListOrders)
			orders.GET("/:orderId", controllers.GetOrder)
			orders.GET("/:orderId/details", controllers.GetOrderWithCustomerDetails)
			orders.PATCH("/:orderId", controllers.UpdateOrder)
			orders.DELETE("/:orderId", controllers.DeleteOrder)
		}

	}

}
