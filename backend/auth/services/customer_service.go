package services

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"strings"

	"github.com/Endale2/DRPS/auth/utils" // Ensure utils package has VerifyGoogleIDToken and CreateToken
	customerModels "github.com/Endale2/DRPS/customers/models"
	customerRepo "github.com/Endale2/DRPS/customers/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OptionalProfile holds non-required profile fields at registration.
type OptionalProfile struct {
	FirstName, LastName, Phone, Address, City, State, PostalCode, Country string
}

var (
	otpStore = struct {
		sync.Mutex
		m map[string]otpEntry
	}{m: make(map[string]otpEntry)}
)

type otpEntry struct {
	OTP     string
	Expires time.Time
}

func generateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendCustomerOTP generates and sends an OTP to the given email
func SendCustomerOTP(email string) error {
	otp := generateOTP()
	otpStore.Lock()
	otpStore.m[strings.ToLower(email)] = otpEntry{OTP: otp, Expires: time.Now().Add(5 * time.Minute)}
	otpStore.Unlock()
	fmt.Printf("[DEBUG] OTP for %s: %s\n", email, otp)
	return nil
}

// VerifyCustomerOTP checks the OTP, creates/fetches customer, returns tokens and profile completeness
func VerifyCustomerOTP(email, otp string) (*customerModels.Customer, string, string, bool, error) {
	emailKey := strings.ToLower(email)
	otpStore.Lock()
	entry, ok := otpStore.m[emailKey]
	otpStore.Unlock()
	if !ok || entry.Expires.Before(time.Now()) || entry.OTP != otp {
		return nil, "", "", false, errors.New("invalid or expired OTP")
	}
	// Remove OTP after use
	otpStore.Lock()
	delete(otpStore.m, emailKey)
	otpStore.Unlock()
	// Find or create customer
	cust, err := customerRepo.GetCustomerByEmail(email)
	if err != nil {
		return nil, "", "", false, err
	}
	if cust == nil {
		now := time.Now()
		cust = &customerModels.Customer{
			Email:     email,
			CreatedAt: now,
			UpdatedAt: now,
		}
		res, err := customerRepo.CreateCustomer(cust)
		if err != nil {
			return nil, "", "", false, err
		}
		cust.ID = res.InsertedID.(primitive.ObjectID)
	}
	// Check profile completeness
	complete := cust.FirstName != "" && cust.LastName != "" && cust.Phone != "" && cust.Address != "" && cust.City != "" && cust.State != "" && cust.PostalCode != "" && cust.Country != ""
	// Generate tokens
	at, err := utils.CreateToken(cust.ID.Hex(), 5*time.Minute)
	if err != nil {
		return nil, "", "", complete, errors.New("failed to generate access token")
	}
	rt, err := utils.CreateToken(cust.ID.Hex(), 7*24*time.Hour)
	if err != nil {
		return nil, "", "", complete, errors.New("failed to generate refresh token")
	}
	return cust, at, rt, complete, nil
}
