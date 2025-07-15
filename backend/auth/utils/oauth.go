package utils

import (
	"context"
	"errors"
	"fmt"
	"os"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
var googleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
var googleRedirectCustomer = os.Getenv("GOOGLE_CUSTOMER_REDIRECT_URI")
var googleRedirectSeller = os.Getenv("GOOGLE_SELLER_REDIRECT_URI")

var telegramCustomerBotToken = os.Getenv("TELEGRAM_CUSTOMER_BOT_TOKEN")
var telegramSellerBotToken = os.Getenv("TELEGRAM_SELLER_BOT_TOKEN")

// GetGoogleOAuthConfigForCustomer returns the Google OAuth config for customers
func GetGoogleOAuthConfigForCustomer() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		RedirectURL:  googleRedirectCustomer,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// GetGoogleOAuthConfigForSeller returns the Google OAuth config for sellers
func GetGoogleOAuthConfigForSeller() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		RedirectURL:  googleRedirectSeller,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// VerifyGoogleIDToken validates the ID token and returns standard claims.
func VerifyGoogleIDToken(token string) (*GooglePayload, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, token, googleClientID)
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

// GetTelegramBotTokenForCustomer returns the Telegram bot token for customers
func GetTelegramBotTokenForCustomer() string {
	return telegramCustomerBotToken
}

// GetTelegramBotTokenForSeller returns the Telegram bot token for sellers
func GetTelegramBotTokenForSeller() string {
	return telegramSellerBotToken
}

// TelegramPayload holds the verified Telegram user info
// See https://core.telegram.org/widgets/login#checking-authorization
type TelegramPayload struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url"`
	AuthDate  int64  `json:"auth_date"`
}

// VerifyTelegramLogin verifies the Telegram login signature and returns the payload if valid
func VerifyTelegramLogin(data map[string]string, botToken string) (*TelegramPayload, error) {
	// 1. Extract the hash
	hash, ok := data["hash"]
	if !ok {
		return nil, errors.New("missing hash in Telegram data")
	}
	// 2. Remove hash from data
	delete(data, "hash")
	// 3. Sort keys and build data_check_string
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var dataCheckString string
	for i, k := range keys {
		if i > 0 {
			dataCheckString += "\n"
		}
		dataCheckString += k + "=" + data[k]
	}
	// 4. Compute HMAC-SHA256 of data_check_string using SHA256(botToken)
	h := sha256.New()
	h.Write([]byte(botToken))
	secret := h.Sum(nil)
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(dataCheckString))
	computedHash := hex.EncodeToString(mac.Sum(nil))
	// 5. Compare hashes (case-insensitive)
	if !hmac.Equal([]byte(computedHash), []byte(hash)) {
		return nil, errors.New("invalid Telegram login signature")
	}
	// 6. Parse payload
	payload := &TelegramPayload{}
	if v, ok := data["id"]; ok {
		// Telegram sends id as string
		var id int64
		fmt.Sscanf(v, "%d", &id)
		payload.ID = id
	}
	payload.FirstName = data["first_name"]
	payload.LastName = data["last_name"]
	payload.Username = data["username"]
	payload.PhotoURL = data["photo_url"]
	if v, ok := data["auth_date"]; ok {
		var ad int64
		fmt.Sscanf(v, "%d", &ad)
		payload.AuthDate = ad
	}
	return payload, nil
}
