package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// jwtSecret is the HMAC signing key for tokens (use environment variable in production)
var jwtSecret = []byte("your-very-secret-key")

// Claims defines the JWT payload structure
// It embeds jwt.RegisteredClaims for standard fields (exp, iat, etc.)
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// CreateJWT creates a signed JWT token and returns it with its duration
func CreateJWT(userID string, duration time.Duration) (string, time.Duration, error) {
	expiresAt := time.Now().Add(duration)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", 0, err
	}
	return signedToken, duration, nil
}

// VerifyJWT parses and validates a JWT token string, returning Claims on success
func VerifyJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
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
