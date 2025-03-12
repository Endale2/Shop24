package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Shop Basic Info
type Shop struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Name           string               `bson:"name"`
	OwnerID        primitive.ObjectID   `bson:"ownerId"` // Reference to user ID
	Description    string               `bson:"description"`
	Products       []primitive.ObjectID `bson:"products"` // Product IDs linked to this shop
	Orders         []primitive.ObjectID `bson:"orders"`   // Order IDs linked to this shop
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
	ShopSettings   ShopSettings         `bson:"shopSettings"`
	ShopFeatures   ShopFeatures         `bson:"shopFeatures"`
	ShopAnalytics  ShopAnalytics        `bson:"shopAnalytics"`
	ShopCustomization ShopCustomization `bson:"shopCustomization"`
}



// Shop Features (Themes, Discounts, and Custom Features)
type ShopFeatures struct {
	Theme         string   `bson:"theme"`        // Shop's theme (e.g., "Modern", "Minimalist")
	Discounts     []Discount `bson:"discounts"`    // Discounts available in the shop
	AffiliateLinks []string `bson:"affiliateLinks"` // Affiliate links used in the shop
	StorePolicies StorePolicies `bson:"storePolicies"` // Store policies (e.g., shipping policy, privacy policy)
}

// Discount struct for discounts available in the shop
type Discount struct {
	Code          string  `bson:"code"`           // Discount code (e.g., "SALE20")
	DiscountRate  float64 `bson:"discountRate"`   // Discount rate (e.g., 0.20 for 20% off)
	ValidUntil    time.Time `bson:"validUntil"`    // Expiry date of the discount
}

// Store Policies (Return, Privacy, Terms, etc.)
type StorePolicies struct {
	PrivacyPolicy string `bson:"privacyPolicy"` // URL or text for privacy policy
	TermsOfService string `bson:"termsOfService"` // URL or text for terms of service
	ReturnPolicy   string `bson:"returnPolicy"`  // Return policy description
}

// Shop Analytics (Sales, Traffic, etc.)
type ShopAnalytics struct {
	TotalSales     float64 `bson:"totalSales"`    // Total sales made by the shop
	TotalOrders    int     `bson:"totalOrders"`   // Total number of orders placed
	TotalVisitors  int     `bson:"totalVisitors"` // Total number of shop visitors
	SalesLast30Days float64 `bson:"salesLast30Days"` // Sales in the last 30 days
	TopSellingProducts []primitive.ObjectID `bson:"topSellingProducts"` // IDs of top-selling products
}

// Shop Customization (Logo, Colors, etc.)
type ShopCustomization struct {
	Logo          string   `bson:"logo"`          // Path to logo image
	PrimaryColor  string   `bson:"primaryColor"`  // Primary color for the shop (e.g., "#FF5733")
	SecondaryColor string  `bson:"secondaryColor"` // Secondary color for the shop (e.g., "#C70039")
	BannerImage   string   `bson:"bannerImage"`   // Path to shop banner image
	FooterText    string   `bson:"footerText"`    // Custom footer text for the shop
}

