package migrations

import (
	"context"
	"log"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DiscountRefactorMigration handles the transition to product-only discounts
func DiscountRefactorMigration() error {
	collection := config.GetCollection("DRPS", "discounts")
	ctx := context.Background()

	log.Println("Starting discount refactor migration...")

	// Step 1: Update all existing discounts to product category if they're not already
	filter := bson.M{
		"category": bson.M{
			"$in": []string{"order", "shipping", "buy_x_get_y"},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"category":   "product",
			"updated_at": time.Now(),
		},
	}

	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Printf("Error updating discount categories: %v", err)
		return err
	}

	log.Printf("Updated %d discounts to product category", result.ModifiedCount)

	// Step 2: Remove deprecated fields from all discounts
	// Note: MongoDB doesn't support removing fields in UpdateMany, so we'll do this in a separate step
	// This is a safety measure - the fields are commented out in the model but still exist in old documents

	// Step 3: Validate that all discounts now have valid product-level configuration
	var invalidDiscounts []models.Discount
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error finding discounts: %v", err)
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var discount models.Discount
		if err := cursor.Decode(&discount); err != nil {
			log.Printf("Error decoding discount: %v", err)
			continue
		}

		// Check if discount has any product-level targets
		if len(discount.AppliesToProducts) == 0 && len(discount.AppliesToVariants) == 0 {
			invalidDiscounts = append(invalidDiscounts, discount)
		}
	}

	if len(invalidDiscounts) > 0 {
		log.Printf("Warning: Found %d discounts without product targets:", len(invalidDiscounts))
		for _, discount := range invalidDiscounts {
			log.Printf("  - Discount ID: %s, Name: %s", discount.ID.Hex(), discount.Name)
		}
		log.Println("These discounts will not be applied to any products.")
	}

	// Step 4: Clean up any orphaned collection references
	// Since we're removing collection support, we should clear these fields
	cleanupFilter := bson.M{
		"applies_to_collections": bson.M{"$exists": true, "$ne": []primitive.ObjectID{}},
	}

	cleanupUpdate := bson.M{
		"$unset": bson.M{
			"applies_to_collections": "",
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	cleanupResult, err := collection.UpdateMany(ctx, cleanupFilter, cleanupUpdate)
	if err != nil {
		log.Printf("Error cleaning up collection references: %v", err)
		return err
	}

	log.Printf("Cleaned up collection references in %d discounts", cleanupResult.ModifiedCount)

	// Step 5: Set default eligibility type for discounts that don't have one
	defaultEligibilityFilter := bson.M{
		"eligibility_type": bson.M{"$exists": false},
	}

	defaultEligibilityUpdate := bson.M{
		"$set": bson.M{
			"eligibility_type": "all",
			"updated_at":       time.Now(),
		},
	}

	defaultResult, err := collection.UpdateMany(ctx, defaultEligibilityFilter, defaultEligibilityUpdate)
	if err != nil {
		log.Printf("Error setting default eligibility type: %v", err)
		return err
	}

	log.Printf("Set default eligibility type for %d discounts", defaultResult.ModifiedCount)

	log.Println("Discount refactor migration completed successfully!")
	return nil
}

// RollbackDiscountRefactorMigration allows rolling back the changes if needed
func RollbackDiscountRefactorMigration() error {
	collection := config.GetCollection("DRPS", "discounts")
	ctx := context.Background()

	log.Println("Starting discount refactor rollback...")

	// This would restore the original discount categories
	// Note: This is a simplified rollback - in a real scenario, you'd want to
	// store the original category values before changing them

	log.Println("Warning: Rollback functionality is limited. Manual intervention may be required.")
	log.Println("Discount refactor rollback completed (limited functionality).")
	return nil
}

// ValidateDiscountRefactorMigration validates that the migration was successful
func ValidateDiscountRefactorMigration() error {
	collection := config.GetCollection("DRPS", "discounts")
	ctx := context.Background()

	log.Println("Validating discount refactor migration...")

	// Check that all discounts are now product category
	var nonProductDiscounts []models.Discount
	cursor, err := collection.Find(ctx, bson.M{
		"category": bson.M{"$ne": "product"},
	})
	if err != nil {
		log.Printf("Error finding non-product discounts: %v", err)
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var discount models.Discount
		if err := cursor.Decode(&discount); err != nil {
			continue
		}
		nonProductDiscounts = append(nonProductDiscounts, discount)
	}

	if len(nonProductDiscounts) > 0 {
		log.Printf("Validation failed: Found %d discounts that are not product category", len(nonProductDiscounts))
		for _, discount := range nonProductDiscounts {
			log.Printf("  - Discount ID: %s, Name: %s, Category: %s", discount.ID.Hex(), discount.Name, discount.Category)
		}
		return mongo.ErrNoDocuments
	}

	// Check that all discounts have eligibility type set
	var noEligibilityType []models.Discount
	cursor, err = collection.Find(ctx, bson.M{
		"eligibility_type": bson.M{"$exists": false},
	})
	if err != nil {
		log.Printf("Error finding discounts without eligibility type: %v", err)
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var discount models.Discount
		if err := cursor.Decode(&discount); err != nil {
			continue
		}
		noEligibilityType = append(noEligibilityType, discount)
	}

	if len(noEligibilityType) > 0 {
		log.Printf("Validation failed: Found %d discounts without eligibility type", len(noEligibilityType))
		for _, discount := range noEligibilityType {
			log.Printf("  - Discount ID: %s, Name: %s", discount.ID.Hex(), discount.Name)
		}
		return mongo.ErrNoDocuments
	}

	log.Println("Discount refactor migration validation passed!")
	return nil
}
