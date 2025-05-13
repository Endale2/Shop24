package utils

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v4"
)

// jwtKey holds the secret key loaded from the JWT_SECRET env var
var jwtKey []byte

func init() {
	// Load environment variables from .env (if present)
	_ = godotenv.Load()

	// Retrieve JWT_SECRET
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is required")
	}
	jwtKey = []byte(secret)
}

// Claims represents the JWT claims payload.
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// CreateToken generates a JWT token string with the given userID and expiry duration.
func CreateToken(userID string, duration time.Duration) (string, error) {
	expiresAt := time.Now().Add(duration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken verifies the JWT token string and returns the Claims if valid.
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}