package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/config"
	authRoutes "github.com/Endale2/DRPS/auth/routes"   // Rename auth routes
	adminRoutes "github.com/Endale2/DRPS/admin/routes" // Rename admin routes
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Create a new Gin router
	r := gin.Default()

	// Configure CORS
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
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

	// Test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MongoDB connected!"})
	})

	// Start the server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
