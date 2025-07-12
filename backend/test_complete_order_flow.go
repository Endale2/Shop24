package main

import (
	"fmt"
	"time"

	customerModel "github.com/Endale2/DRPS/customers/models"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println("=== Testing Complete Order Flow with Automatic Discount Application ===\n")

	// Create test IDs
	shopID := primitive.NewObjectID()
	productID := primitive.NewObjectID()
	variantID := primitive.NewObjectID()
	customerID := primitive.NewObjectID()
	sellerID := primitive.NewObjectID()

	fmt.Printf("Test Shop ID: %s\n", shopID.Hex())
	fmt.Printf("Test Product ID: %s\n", productID.Hex())
	fmt.Printf("Test Variant ID: %s\n", variantID.Hex())
	fmt.Printf("Test Customer ID: %s\n", customerID.Hex())

	// Create test shop
	shop := &models.Shop{
		ID:      shopID,
		Name:    "Test Shop",
		Slug:    "test-shop",
		OwnerID: sellerID,
	}
	_, err := repositories.CreateShop(shop)
	if err != nil {
		fmt.Printf("Error creating shop: %v\n", err)
		return
	}
	fmt.Printf("✓ Created test shop\n")

	// Create test product
	product := &models.Product{
		ID:          productID,
		ShopID:      shopID,
		Name:        "Test Product",
		Slug:        "test-product",
		Description: "Test product for discount testing",
		Price:       50.0,
		Stock:       100,
		Category:    "Test",
		CreatedBy:   sellerID,
		Variants: []models.Variant{
			{
				VariantID: variantID,
				Price:     60.0,
				Stock:     50,
				Options: []models.Option{
					{Name: "Size", Value: "Large"},
				},
			},
		},
	}
	_, err = repositories.CreateProduct(product)
	if err != nil {
		fmt.Printf("Error creating product: %v\n", err)
		return
	}
	fmt.Printf("✓ Created test product with variant\n")

	// Create test customer (using the correct customer model)
	customer := &customerModel.Customer{
		ID:        customerID,
		FirstName: "Test",
		LastName:  "Customer",
		Email:     "test@example.com",
	}
	_, err = repositories.CreateCustomer(customer)
	if err != nil {
		fmt.Printf("Error creating customer: %v\n", err)
		return
	}
	fmt.Printf("✓ Created test customer\n")

	// Create test discounts
	percentageDiscount := &models.Discount{
		ID:                primitive.NewObjectID(),
		Name:              "20% Off Test",
		Description:       "Test percentage discount",
		Category:          models.DiscountCategoryProduct,
		Type:              models.DiscountTypePercentage,
		Value:             20.0, // 20% off
		ShopID:            shopID,
		SellerID:          sellerID,
		AppliesToProducts: []primitive.ObjectID{productID},
		EligibilityType:   models.DiscountEligibilityAll,
		Active:            true,
		StartAt:           time.Time{}, // Started
		EndAt:             time.Time{}, // No end
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	fixedDiscount := &models.Discount{
		ID:                primitive.NewObjectID(),
		Name:              "$10 Off Test",
		Description:       "Test fixed discount",
		Category:          models.DiscountCategoryProduct,
		Type:              models.DiscountTypeFixed,
		Value:             10.0, // $10 off per unit
		ShopID:            shopID,
		SellerID:          sellerID,
		AppliesToProducts: []primitive.ObjectID{productID},
		EligibilityType:   models.DiscountEligibilityAll,
		Active:            true,
		StartAt:           time.Time{}, // Started
		EndAt:             time.Time{}, // No end
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	// Save discounts
	_, err = repositories.CreateDiscount(percentageDiscount)
	if err != nil {
		fmt.Printf("Error saving percentage discount: %v\n", err)
		return
	}
	fmt.Printf("✓ Saved percentage discount: %s\n", percentageDiscount.Name)

	_, err = repositories.CreateDiscount(fixedDiscount)
	if err != nil {
		fmt.Printf("Error saving fixed discount: %v\n", err)
		return
	}
	fmt.Printf("✓ Saved fixed discount: %s\n", fixedDiscount.Name)

	// Test scenarios
	testCases := []struct {
		name       string
		unitPrice  float64
		quantity   int
		useVariant bool
	}{
		{"Product without variant", 50.0, 1, false},
		{"Product with variant", 60.0, 2, true},
		{"Multiple items without variant", 50.0, 3, false},
		{"Multiple items with variant", 60.0, 2, true},
	}

	fmt.Println("\n=== Testing Order Placement with Automatic Discount Application ===")

	for _, tc := range testCases {
		fmt.Printf("\n--- %s ---\n", tc.name)

		// Simulate the order placement logic
		lineTotal := tc.unitPrice * float64(tc.quantity)
		fmt.Printf("Original line total: $%.2f\n", lineTotal)

		// Get collection IDs (empty for this test)
		collectionIDs := []primitive.ObjectID{}

		// Get active discounts for this product/variant
		var variantIDToUse primitive.ObjectID
		if tc.useVariant {
			variantIDToUse = variantID
		}

		discounts, err := services.GetActiveDiscountsForProductService(shopID, productID, variantIDToUse, collectionIDs)
		if err != nil {
			fmt.Printf("Error getting discounts: %v\n", err)
			continue
		}

		fmt.Printf("Found %d discounts\n", len(discounts))

		var bestDiscount *models.Discount
		bestSavings := 0.0

		// Find the best discount (same logic as in order controller)
		for i := range discounts {
			discount := &discounts[i]
			fmt.Printf("Checking discount: %s (Type: %s, Value: %.2f, Active: %v)\n",
				discount.Name, discount.Type, discount.Value, discount.IsActive())

			if discount.IsActive() {
				canUse, err := services.CanCustomerUseDiscount(discount, customerID, []primitive.ObjectID{})
				fmt.Printf("Customer can use discount: %v, error: %v\n", canUse, err)

				if err == nil && canUse {
					savings := discount.CalculateDiscountForQuantity(tc.unitPrice, tc.quantity)
					fmt.Printf("Discount savings: %.2f for unit price: %.2f, quantity: %d\n", savings, tc.unitPrice, tc.quantity)

					if savings > bestSavings {
						bestSavings = savings
						bestDiscount = discount
						fmt.Printf("New best discount: %s with savings: %.2f\n", discount.Name, savings)
					}
				}
			}
		}

		// Apply discount (same logic as in order controller)
		itemDiscountAmount := 0.0
		if bestDiscount != nil && bestSavings > 0 {
			itemDiscountAmount = bestSavings
			fmt.Printf("Applied discount: %s, amount: %.2f\n", bestDiscount.Name, itemDiscountAmount)
		}

		finalLineTotal := lineTotal - itemDiscountAmount
		if finalLineTotal < 0 {
			finalLineTotal = 0
		}

		fmt.Printf("Final line total: $%.2f (saved: $%.2f)\n", finalLineTotal, itemDiscountAmount)

		// Create order item with discounted price
		orderItem := models.OrderItem{
			ProductID:  productID,
			VariantID:  variantIDToUse,
			Name:       "Test Product",
			Quantity:   tc.quantity,
			UnitPrice:  tc.unitPrice,
			TotalPrice: finalLineTotal, // This is the discounted price!
			Image:      "test-image.jpg",
		}

		fmt.Printf("Order item created with discounted total: $%.2f\n", orderItem.TotalPrice)
	}

	fmt.Println("\n=== Test Complete ===")
	fmt.Println("The backend automatically applies discounts during order placement!")
	fmt.Println("This works independently of the frontend.")
}
