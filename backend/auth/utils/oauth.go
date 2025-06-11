package utils

import (
    "context"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/idtoken"
)

var GoogleOAuthConfig *oauth2.Config

func init() {
    GoogleOAuthConfig = &oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"), // e.g. http://localhost:8080/auth/seller/oauth/google/callback
        Scopes:       []string{"openid", "email", "profile"},
        Endpoint:     google.Endpoint,
    }
}

// VerifyGoogleIDToken validates the ID token and returns standard claims.
func VerifyGoogleIDToken(token string) (*GooglePayload, error) {
    ctx := context.Background()
    payload, err := idtoken.Validate(ctx, token, GoogleOAuthConfig.ClientID)
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

type GooglePayload struct {
    Sub, Email, GivenName, FamilyName, Picture string
}
