package middlewares

import (
	"net/http"

	"github.com/Endale2/DRPS/auth/utils"
	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware ensures access_token is valid and extracts the admin ID.
// auth/middleware/auth_middleware.go

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		// store the hex userID string
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
