package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Endale2/DRPS/config"
	authRoutes "github.com/Endale2/DRPS/auth/routes"
	adminRoutes "github.com/Endale2/DRPS/admin/routes"
	sellerRoutes "github.com/Endale2/DRPS/sellers/routes"
	storefrontRoutes "github.com/Endale2/DRPS/customers/routes"
)

func main() {
	// Load environment variables from .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	// Check for required JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("❌ JWT_SECRET must be set")
	}

	// Connect to MongoDB or other DB
	config.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// CORS configuration: only allow dashboard and root storefront
	corsConfig := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", // Seller Dashboard
			"http://localhost:5174", // Storefront root
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply CORS middleware
	r.Use(cors.New(corsConfig))

	// Register routes
	authRoutes.AuthRoutes(r)
	adminRoutes.AdminRoutes(r)
	sellerRoutes.SellerRoute(r)
	storefrontRoutes.StorefrontRoutes(r)

	// Health check route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MongoDB connected!"})
	})

	// Start server
	log.Println("🚀 Server running on port 8080")
	r.Run(":8080")
}
