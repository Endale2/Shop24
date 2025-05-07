package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Secret key for signing JWTs (use env var or config in production)
var jwtKey = []byte("your_secret_key")

// Custom claims structure
type Claims struct {
	AdminID string `json:"admin_id"`
	jwt.RegisteredClaims
}

// ValidateToken checks if the token is valid and returns the claims
func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

// AuthMiddleware is a Gin middleware that protects routes requiring authentication
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from HTTP-only cookie
		token, err := c.Cookie("access_token")
		if err != nil || strings.TrimSpace(token) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - token missing"})
			return
		}

		// Validate the token
		claims, err := ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - invalid token"})
			return
		}

		// Store the admin ID in context for access in handlers
		c.Set("admin_id", claims.AdminID)
		c.Next()
	}
}
