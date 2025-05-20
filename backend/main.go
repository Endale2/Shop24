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
)

func main() {
	// ✅ Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found or failed to load. Using system environment variables.")
	}

	// Example: Access env variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("❌ JWT_SECRET is not set in .env or environment")
	}

	// Connect to MongoDB
	config.ConnectDB()

	// Create a new Gin router
	r := gin.Default()

	// Configure CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	// Routes
	authRoutes.AuthRoutes(r)
	adminRoutes.AdminRoutes(r)
	sellerRoutes.SellerRoute(r)

	// Test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MongoDB connected!"})
	})

	// Start the server
	log.Println("🚀 Server running on port 8080")
	r.Run(":8080")
}
