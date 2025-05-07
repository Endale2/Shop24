package middleware

import (
	"net/http"

	"github.com/Endale2/DRPS/auth/services"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for the access_token in the HTTP-only cookie
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read access token from HTTP-only cookie
		accessToken, err := c.Cookie("access_token")
		if err != nil || accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access token not found in cookies"})
			return
		}

		// Validate the JWT access token
		claims, err := services.ValidateToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired access token"})
			return
		}

		// Set relevant info (like admin ID) in context for use in handlers
		c.Set("admin_id", claims.UserID) // or claims.AdminID depending on naming
		c.Next()
	}
}
