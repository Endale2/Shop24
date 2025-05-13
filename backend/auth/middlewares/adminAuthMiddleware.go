package middlewares

import (
	"net/http"
	"strings"

	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware ensures access_token is valid and extracts the admin ID.
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from HTTP-only cookie
		token, err := c.Cookie("access_token")
		if err != nil || strings.TrimSpace(token) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - token missing"})
			return
		}

		// Validate the token using utils.ParseToken
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - invalid or expired token"})
			return
		}

		// Set admin_id into context
		c.Set("admin_id", claims.UserID) // UserID in claims maps to AdminID
		c.Next()
	}
}
