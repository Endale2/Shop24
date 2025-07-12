# Discount System Refactor

## Overview

This refactor simplifies the discount system to support **only product-level discounts** (fixed or percentage), removing order-level, shipping, and buy-x-get-y discount functionality. The system maintains full customer eligibility and usage tracking capabilities.

## Changes Made

### 1. Model Changes (`backend/shared/models/discount.go`)

**Removed:**
- `DiscountCategoryOrder` - order-level discounts
- `DiscountCategoryShipping` - shipping/free shipping discounts  
- `DiscountCategoryBuyXGetY` - buy X get Y discounts
- `AppliesToCollections` - collection-level targeting
- `BuyProductIDs`, `BuyQuantity`, `GetProductIDs`, `GetQuantity` - buy X get Y fields
- `MinimumOrderSubtotal` - order-level minimum requirements
- `FreeShipping`, `MinimumOrderForFreeShipping` - shipping-level fields

**Kept:**
- `DiscountCategoryProduct` - product-level discounts only
- `AppliesToProducts` - specific product targeting
- `AppliesToVariants` - specific variant targeting
- All eligibility fields (`EligibilityType`, `AllowedCustomerIDs`, `AllowedSegmentIDs`)
- All usage tracking fields (`UsageLimit`, `PerCustomerLimit`, `CurrentUsage`, `UsageTracking`)
- All timing fields (`StartAt`, `EndAt`, `Active`)

### 2. Repository Changes (`backend/shared/repositories/discount_repository.go`)

**Removed:**
- `GetActiveOrderDiscountsForShop()` - order discount queries
- `GetActiveShippingDiscountsForShop()` - shipping discount queries

**Updated:**
- `GetActiveDiscountsForProduct()` - removed collection logic, kept for future reintroduction

### 3. Service Changes (`backend/shared/services/discount_service.go`)

**Removed:**
- `GetActiveOrderDiscountsForShopService()` - order discount service
- `GetActiveShippingDiscountsForShopService()` - shipping discount service
- `ApplyDiscountsToCart()` - order-level cart discount application

**Updated:**
- `ApplyDiscountsToProduct()` - product-level discount application only
- `GetBestProductDiscount()` - product-level discount selection only

### 4. Cart Service Changes (`backend/shared/services/cart_service.go`)

**Removed:**
- `OrderDiscountDetail` struct - order-level discount details
- Order-level discount application logic

**Updated:**
- `CartWithDiscountDetails` - only includes item-level discount details
- `CalculateTotals()` - only applies product-level discounts

### 5. Controller Changes (`backend/sellers/controllers/discount_controller.go`)

**Removed:**
- All order-level, shipping, and buy-x-get-y field handling
- Collection-level discount field processing

**Updated:**
- `DiscountInput` struct - removed deprecated fields
- All CRUD operations - simplified to product-level only
- Response formatting - removed deprecated fields

## Database Migration

A migration script is provided at `backend/migrations/discount_refactor_migration.go` that:

1. **Updates existing discounts** to product category
2. **Cleans up deprecated fields** from the database
3. **Validates the migration** was successful
4. **Sets default eligibility** for discounts without it

### Running the Migration

```go
import "github.com/Endale2/DRPS/migrations"

// Run the migration
err := migrations.DiscountRefactorMigration()
if err != nil {
    log.Fatal("Migration failed:", err)
}

// Validate the migration
err = migrations.ValidateDiscountRefactorMigration()
if err != nil {
    log.Fatal("Validation failed:", err)
}
```

## API Changes

### Creating Discounts

**Before:**
```json
{
  "name": "Summer Sale",
  "category": "order",
  "type": "percentage",
  "value": 10,
  "minimumOrderSubtotal": 50.0,
  "freeShipping": true,
  "appliesToCollections": ["collection1", "collection2"]
}
```

**After:**
```json
{
  "name": "Summer Sale",
  "category": "product",
  "type": "percentage", 
  "value": 10,
  "appliesToProducts": ["product1", "product2"],
  "appliesToVariants": ["variant1", "variant2"],
  "eligibilityType": "all"
}
```

### Discount Response

**Before:**
```json
{
  "id": "discount_id",
  "name": "Summer Sale",
  "category": "order",
  "free_shipping": true,
  "minimum_order_subtotal": 50.0,
  "applies_to_collections": ["collection1"],
  "buy_product_ids": ["product1"],
  "get_product_ids": ["product2"]
}
```

**After:**
```json
{
  "id": "discount_id", 
  "name": "Summer Sale",
  "category": "product",
  "type": "percentage",
  "value": 10,
  "applies_to_products": ["product1", "product2"],
  "applies_to_variants": ["variant1"],
  "eligibility_type": "all",
  "usage_limit": 100,
  "per_customer_limit": 1
}
```

## Eligibility System

The eligibility system remains fully functional:

### Eligibility Types

1. **`all`** - Everyone can use (default)
2. **`specific`** - Only customers in `allowed_customers` list
3. **`segment`** - Only customers in segments listed in `allowed_segments`

### Managing Eligibility

**Add customers to discount:**
```http
POST /seller/shops/{shopId}/discounts/{discountId}/customers
{
  "customerIds": ["customer1", "customer2"]
}
```

**Add segments to discount:**
```http
POST /seller/shops/{shopId}/discounts/{discountId}/segments
{
  "segmentIds": ["segment1", "segment2"]
}
```

**Validate customer eligibility:**
```http
POST /seller/shops/{shopId}/discounts/{discountId}/validate
{
  "customerId": "customer1"
}
```

## Usage Tracking

Usage tracking remains fully functional:

- **Overall usage limits** - `usage_limit` field
- **Per-customer limits** - `per_customer_limit` field  
- **Usage statistics** - Available via `/discounts/{id}/usage` endpoint
- **Automatic tracking** - Usage is recorded when discounts are applied

## Multiple Discounts

Multiple discounts can apply to the same product/variant. The system:

1. **Finds all applicable discounts** for a product/variant
2. **Validates customer eligibility** for each discount
3. **Selects the best discount** (highest value) that the customer is eligible for
4. **Applies the discount** and records usage

## Backward Compatibility

### What Still Works

- ✅ All existing product-level discounts continue to work
- ✅ Customer eligibility system remains unchanged
- ✅ Usage tracking and limits remain functional
- ✅ All CRUD operations for discounts
- ✅ Discount validation and application logic

### What's Changed

- ❌ Order-level discounts no longer work
- ❌ Shipping discounts no longer work  
- ❌ Buy X Get Y discounts no longer work
- ❌ Collection-level targeting no longer works
- ❌ Cart-level discount application removed

### Migration Notes

1. **Existing discounts** are automatically converted to product category
2. **Deprecated fields** are preserved in database but ignored by new code
3. **Collection references** are cleaned up during migration
4. **Default eligibility** is set to "all" for existing discounts

## Future Reintroduction

All removed functionality is marked with `TODO` comments for easy reintroduction:

```go
// TODO: Reintroduce order-level functionality in future versions
// TODO: Reintroduce shipping-level functionality in future versions  
// TODO: Reintroduce Buy X Get Y functionality in future versions
// TODO: Reintroduce collection-level discounts in future versions
```

## Testing

### Unit Tests

Test the new discount functionality:

```go
// Test product-level discount creation
discount := &models.Discount{
    Name: "Test Discount",
    Category: models.DiscountCategoryProduct,
    Type: models.DiscountTypePercentage,
    Value: 10,
    AppliesToProducts: []primitive.ObjectID{productID},
    EligibilityType: models.DiscountEligibilityAll,
}

// Test eligibility validation
canUse, err := CanCustomerUseDiscount(discount, customerID, customerSegmentIDs)

// Test discount calculation
amount := discount.CalculateDiscount(100.0) // Returns 10.0
```

### Integration Tests

Test the API endpoints:

```bash
# Create a product-level discount
curl -X POST /seller/shops/{shopId}/discounts \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Discount",
    "category": "product", 
    "type": "percentage",
    "value": 10,
    "appliesToProducts": ["product1"],
    "startAt": "2024-01-01T00:00:00Z",
    "endAt": "2024-12-31T23:59:59Z",
    "active": true
  }'

# Validate customer eligibility
curl -X POST /seller/shops/{shopId}/discounts/{discountId}/validate \
  -H "Content-Type: application/json" \
  -d '{"customerId": "customer1"}'
```

## Performance Impact

The refactor improves performance by:

1. **Simplified queries** - No more complex order/shipping discount lookups
2. **Reduced data transfer** - Smaller discount objects
3. **Faster validation** - Simpler eligibility checks
4. **Optimized cart calculations** - Product-level only

## Monitoring

Monitor the migration and new system:

1. **Migration logs** - Check for any conversion issues
2. **Discount usage** - Monitor that discounts are being applied correctly
3. **Customer eligibility** - Verify eligibility validation works
4. **Performance metrics** - Compare before/after performance

## Rollback Plan

If issues arise, the system can be rolled back by:

1. **Restoring old model fields** - Uncomment deprecated fields
2. **Restoring old services** - Uncomment removed service methods
3. **Database rollback** - Use the rollback migration (limited functionality)
4. **API compatibility** - Restore old API response formats

## Support

For questions or issues with the discount refactor:

1. Check the migration logs for any conversion errors
2. Verify discount eligibility using the validation endpoint
3. Test discount application with the cart service
4. Review the TODO comments for reintroduction guidance 