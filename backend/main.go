package main

import (
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	adminRoutes "github.com/Endale2/DRPS/admin/routes"
	authRoutes "github.com/Endale2/DRPS/auth/routes"
	"github.com/Endale2/DRPS/config"
	storefrontRoutes "github.com/Endale2/DRPS/customers/routes"
	sellerRoutes "github.com/Endale2/DRPS/sellers/routes"
)

func main() {
	// Load environment variables from .env if present (must be first!)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	// Warn if GOOGLE_CUSTOMER_CLIENT_ID is missing
	if os.Getenv("GOOGLE_CUSTOMER_CLIENT_ID") == "" {
		log.Println("❌ GOOGLE_CUSTOMER_CLIENT_ID is not set! Customer OAuth will fail.")
	}

	// Check for required JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("❌ JWT_SECRET must be set")
	}

	// Connect to MongoDB or other DB
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	// Set Gin to release mode to suppress debug endpoint and warning logs
	gin.SetMode(gin.ReleaseMode)
	// Initialize Gin router
	r := gin.Default()

	// CORS: allow three exact origins plus any subdomain on 5175
	corsConfig := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			u, err := url.Parse(origin)
			if err != nil {
				return false
			}
			host := u.Hostname()
			port := u.Port()

			// 1) Admin Dashboard
			if host == "localhost" && port == "5173" {
				return true
			}
			// 2) Seller root app
			if host == "localhost" && port == "5174" {
				return true
			}
			// 3) Storefront root
			if host == "localhost" && port == "5175" {
				return true
			}
			// 4) Any storefront subdomain on 5175, e.g. shop123.localhost:5175
			if port == "5175" && strings.HasSuffix(host, ".localhost") {
				return true
			}
			return false
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
		if config.DB == nil {
			c.JSON(500, gin.H{"error": "MongoDB not connected"})
			return
		}
		c.JSON(200, gin.H{"message": "MongoDB connected!"})
	})

	// Start server
	log.Println("🚀 Server running on port 8080")
	r.Run(":8080")
}
