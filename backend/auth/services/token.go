package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UserID string `json:"user_id"` // Change to AdminID if that's more accurate
	jwt.RegisteredClaims
}

func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, errors.New("could not parse token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
