package utils

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

// Seller Google OAuth config
var googleSellerClientID = os.Getenv("GOOGLE_SELLER_CLIENT_ID")
var googleSellerClientSecret = os.Getenv("GOOGLE_SELLER_CLIENT_SECRET")
var googleSellerRedirect = os.Getenv("GOOGLE_SELLER_REDIRECT_URI")

// GetGoogleOAuthConfigForSeller returns the Google OAuth config for sellers
func GetGoogleOAuthConfigForSeller() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_SELLER_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_SELLER_CLIENT_SECRET")
	redirect := os.Getenv("GOOGLE_SELLER_REDIRECT_URI")
	if clientID == "" || clientSecret == "" || redirect == "" {
		fmt.Printf("[FATAL] Seller OAuth env missing: client_id='%s', client_secret='%s', redirect_uri='%s'\n", clientID, clientSecret, redirect)
		panic("Missing required seller OAuth environment variables")
	}
	fmt.Println("SELLER CLIENT ID:", clientID)
	fmt.Println("SELLER REDIRECT URI:", redirect)
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirect,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// VerifyGoogleIDToken validates the ID token and returns standard claims.
type GooglePayload struct {
	Sub, Email, GivenName, FamilyName, Picture string
}

func VerifyGoogleIDToken(token string, clientID string) (*GooglePayload, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, token, clientID)
	if err != nil {
		return nil, err
	}
	return &GooglePayload{
		Sub:        payload.Subject,
		Email:      payload.Claims["email"].(string),
		GivenName:  payload.Claims["given_name"].(string),
		FamilyName: payload.Claims["family_name"].(string),
		Picture:    payload.Claims["picture"].(string),
	}, nil
}

// Exported functions to get the correct client IDs for use in services
func GoogleSellerClientID() string {
	return googleSellerClientID
}

// Exported functions to get the correct frontend URLs for use in controllers
func SellerFrontendURL() string {
	return os.Getenv("SELLER_FRONTEND_URL")
}
