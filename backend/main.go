package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/auth/routes" //auth routes
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	// Create a new Gin router
	r := gin.Default()
    //routes
	routes.RegisterAuthRoutes(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MongoDB connected!"})
	})

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
