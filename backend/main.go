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

	"github.com/Endale2/DRPS/config"
	authRoutes   "github.com/Endale2/DRPS/auth/routes"
	adminRoutes  "github.com/Endale2/DRPS/admin/routes"
	sellerRoutes "github.com/Endale2/DRPS/sellers/routes"
	storefrontRoutes "github.com/Endale2/DRPS/customers/routes"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("❌ JWT_SECRET must be set")
	}

	config.ConnectDB()

	r := gin.Default()

	// CORS that allows:
	// - http://localhost:5173       (dashboard)
	// - http://localhost:5174       (storefront root)
	// - http://<any>.localhost:5174 (any storefront subdomain)
	corsConfig := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			u, err := url.Parse(origin)
			if err != nil {
				return false
			}
			host := u.Hostname()
			port := u.Port()

			// Seller  Dashboard front-end
			if host == "localhost" && port == "5173" {
				return true
			}

			// Storefront root
			if host == "localhost" && port == "5174" {
				return true
			}

			// Any subdomain storefront, e.g. shop123.localhost:5174
			if port == "5174" && strings.HasSuffix(host, ".localhost") {
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

	r.Use(cors.New(corsConfig))

	// your existing routes
	authRoutes.AuthRoutes(r)
	adminRoutes.AdminRoutes(r)
	sellerRoutes.SellerRoute(r)
	storefrontRoutes.StorefrontRoutes(r)
	

	// health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "MongoDB connected!"})
	})

	log.Println("🚀 Server running on port 8080")
	r.Run(":8080")
}
