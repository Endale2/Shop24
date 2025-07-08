package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sellerRepos "github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// slugify takes a name (e.g. "Catan Base Game") and returns
// a simple URL‑friendly slug ("catan-base-game").
func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.Join(strings.Fields(s), "-")
	return s
}

// normalizeProduct ensures that:
//  1. If no variants exist, we inject a "default" variant
//  2. Each variant has a non‑zero VariantID, CreatedAt, UpdatedAt
//  3. The Product's DisplayPrice is set to the minimum Variant price
//  4. The Product's MainImage is set from Images[0] or first variant image
//  5. The Product's Stock is the sum of its variant stocks
func normalizeProduct(p *models.Product) {
	now := time.Now()

	// 1) Only create default variant if the product truly has variants in the database
	// Products without variants should remain as simple products
	hasRealVariants := false
	for _, v := range p.Variants {
		// A real variant has meaningful options (not empty name/value)
		if len(v.Options) > 0 && (v.Options[0].Name != "" || v.Options[0].Value != "") {
			hasRealVariants = true
			break
		}
	}

	// 2) Ensure every variant has timestamps (IDs are handled by EnsureProductVariantIDs)
	for i := range p.Variants {
		v := &p.Variants[i]
		if v.CreatedAt.IsZero() {
			v.CreatedAt = now
		}
		v.UpdatedAt = now
	}

	// 3) If product has real variants, calculate aggregate values
	if hasRealVariants {
		minPrice := p.Variants[0].Price
		maxDisplayPrice := p.Variants[0].DisplayPrice
		totalStock := 0
		for _, v := range p.Variants {
			if v.Price < minPrice {
				minPrice = v.Price
			}
			if v.DisplayPrice != nil && (maxDisplayPrice == nil || *v.DisplayPrice > *maxDisplayPrice) {
				maxDisplayPrice = v.DisplayPrice
			}
			totalStock += v.Stock
		}
		p.Price = minPrice
		p.DisplayPrice = maxDisplayPrice
		p.Stock = totalStock
	}
	// If no real variants, keep the product's original price, display_price, and stock

	// 4) Determine MainImage
	if len(p.Images) > 0 {
		p.MainImage = p.Images[0]
	} else {
		for _, v := range p.Variants {
			if v.Image != "" {
				p.MainImage = v.Image
				break
			}
		}
	}
}

// CreateProductService inserts a new Product into MongoDB after applying
// business logic (ID, slug, normalization including MainImage and Stock).
func CreateProductService(p *models.Product) (*mongo.InsertOneResult, error) {
	if strings.TrimSpace(p.Name) == "" {
		return nil, errors.New("product name is required")
	}

	now := time.Now()
	p.ID = primitive.NewObjectID()
	p.CreatedAt = now
	p.UpdatedAt = now

	if strings.TrimSpace(p.Slug) == "" {
		p.Slug = slugify(p.Name)
	}

	normalizeProduct(p)
	return repositories.CreateProduct(p)
}

// GetProductByIDService retrieves a Product by its hex ID, then normalizes it.
func GetProductByIDService(id string) (*models.Product, error) {
	p, err := repositories.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if p != nil {
		// Ensure variant IDs are set and saved to database
		err = EnsureProductVariantIDs(p)
		if err != nil {
			return nil, err
		}
		normalizeProduct(p)
	}
	return p, nil
}

// GetAllProductsService returns all Products. Each is normalized.
func GetAllProductsService() ([]models.Product, error) {
	list, err := repositories.GetAllProducts()
	if err != nil {
		return nil, err
	}
	for i := range list {
		normalizeProduct(&list[i])
	}
	return list, nil
}

// UpdateProductService updates fields by ID. It also handles recalculating
// the total stock if the variants are updated.
func UpdateProductService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	updatedData["updatedAt"] = time.Now()

	if newName, ok := updatedData["name"].(string); ok && strings.TrimSpace(newName) != "" {
		updatedData["slug"] = slugify(newName)
	}

	// If variants are updated, we expect each to include an 'options' array
	if rawVariants, ok := updatedData["variants"].([]interface{}); ok {
		totalStock := 0
		minPrice := 0.0
		maxDisplayPrice := (*float64)(nil)
		for idx, rv := range rawVariants {
			if vMap, isMap := rv.(map[string]interface{}); isMap {
				// Convert the incoming 'options' field to []Option
				if optsRaw, hasOpts := vMap["options"].([]interface{}); hasOpts {
					var opts []models.Option
					for _, o := range optsRaw {
						if oMap, ok := o.(map[string]interface{}); ok {
							name, _ := oMap["name"].(string)
							value, _ := oMap["value"].(string)
							opts = append(opts, models.Option{Name: name, Value: value})
						}
					}
					rawVariants[idx].(map[string]interface{})["options"] = opts
				}
				if priceF, hasPrice := vMap["price"].(float64); hasPrice {
					if idx == 0 || priceF < minPrice {
						minPrice = priceF
					}
				}
				if dispF, hasDisp := vMap["display_price"].(float64); hasDisp {
					if maxDisplayPrice == nil || dispF > *maxDisplayPrice {
						maxDisplayPrice = new(float64)
						*maxDisplayPrice = dispF
					}
				}
				if stockF, hasStock := vMap["stock"].(float64); hasStock {
					totalStock += int(stockF)
				}
			}
		}
		updatedData["price"] = minPrice
		updatedData["display_price"] = maxDisplayPrice
		updatedData["stock"] = totalStock
		// Lock these fields from being edited directly ONLY if variants are present
		if _, exists := updatedData["price"]; exists {
			delete(updatedData, "price")
		}
		if _, exists := updatedData["display_price"]; exists {
			delete(updatedData, "display_price")
		}
		if _, exists := updatedData["stock"]; exists {
			delete(updatedData, "stock")
		}
	}
	// For products with NO variants, allow price, display_price, stock, and main_image to be updated directly
	// No further action needed

	return repositories.UpdateProduct(id, updatedData)
}

// DeleteProductService removes the product document by its ID.
func DeleteProductService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteProduct(id)
}

// CountProductsService returns how many products match a given filter.
func CountProductsService(filter bson.M) (int64, error) {
	return repositories.CountProducts(filter)
}

// CountByCategoryService returns a map[category]count by aggregating in MongoDB.
func CountByCategoryService() (map[string]int64, error) {
	return repositories.CountByCategory()
}

// GetProductsByShopIDService fetches all products whose "shop_id" equals shopOID.
// Each product is normalized before returning.
func GetProductsByShopIDService(shopOID primitive.ObjectID) ([]models.Product, error) {
	list, err := repositories.GetProductsByFilter(bson.M{"shop_id": shopOID})
	if err != nil {
		return nil, err
	}
	for i := range list {
		normalizeProduct(&list[i])
	}
	return list, nil
}

// GetProductsByShopSlugService looks up a Shop by slug, then fetches all products
// belonging to that shop. Each product is normalized before returning.
func GetProductsByShopSlugService(slug string) ([]models.Product, error) {
	list, err := repositories.GetProductsByShopSlug(slug)
	if err != nil {
		return nil, err
	}
	for i := range list {
		// Ensure variant IDs are set and saved to database
		err = EnsureProductVariantIDs(&list[i])
		if err != nil {
			return nil, err
		}
		normalizeProduct(&list[i])
	}
	return list, nil
}

// GetProductBySlugService retrieves a single Product document by its slug,
// then normalizes it (setting display price, etc.).
func GetProductBySlugService(slug string) (*models.Product, error) {
	p, err := repositories.GetProductBySlug(slug)
	if err != nil {
		return nil, err
	}
	if p != nil {
		// Ensure variant IDs are set and saved to database
		err = EnsureProductVariantIDs(p)
		if err != nil {
			return nil, err
		}
		normalizeProduct(p)
	}
	return p, nil
}

// ReduceProductStock reduces stock for a product (no variants)
func ReduceProductStock(productID primitive.ObjectID, quantity int) error {
	// Get current product
	product, err := repositories.GetProductByID(productID.Hex())
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	// Check if we have enough stock
	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	// Update stock
	newStock := product.Stock - quantity
	_, err = repositories.UpdateProduct(productID.Hex(), bson.M{
		"stock":     newStock,
		"updatedAt": time.Now(),
	})
	return err
}

// ReduceVariantStock reduces stock for a specific variant
func ReduceVariantStock(productID, variantID primitive.ObjectID, quantity int) error {
	// Get current product
	product, err := repositories.GetProductByID(productID.Hex())
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	// Find the variant
	var targetVariant *models.Variant
	for i := range product.Variants {
		if product.Variants[i].VariantID == variantID {
			targetVariant = &product.Variants[i]
			break
		}
	}
	if targetVariant == nil {
		return errors.New("variant not found")
	}

	// Check if we have enough stock
	if targetVariant.Stock < quantity {
		return errors.New("insufficient stock for variant")
	}

	// Update variant stock
	targetVariant.Stock -= quantity
	targetVariant.UpdatedAt = time.Now()

	// Recalculate total stock
	totalStock := 0
	for _, v := range product.Variants {
		totalStock += v.Stock
	}

	// Update product with new variant stock and total stock
	_, err = repositories.UpdateProduct(productID.Hex(), bson.M{
		"variants":  product.Variants,
		"stock":     totalStock,
		"updatedAt": time.Now(),
	})
	return err
}

// GetProductByIDWithDiscountsService retrieves a Product by its hex ID, applies active discounts, and normalizes it.
func GetProductByIDWithDiscountsService(id string, customerID *primitive.ObjectID) (*models.Product, error) {
	p, err := repositories.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, nil
	}

	// Ensure variant IDs are set and saved to database
	err = EnsureProductVariantIDs(p)
	if err != nil {
		return nil, err
	}

	normalizeProduct(p)
	// Fetch active discounts for this shop/product/variants
	discounts, _ := GetActiveDiscountsForProductService(p.ShopID, p.ID, primitive.NilObjectID, nil)
	// Apply discounts to product and variants
	ApplyDiscountsToProduct(p, discounts)
	return p, nil
}

// GetAllProductsWithDiscountsService returns all Products with discounts applied.
func GetAllProductsWithDiscountsService(customerID *primitive.ObjectID) ([]models.Product, error) {
	list, err := repositories.GetAllProducts()
	if err != nil {
		return nil, err
	}
	for i := range list {
		normalizeProduct(&list[i])
		discounts, _ := GetActiveDiscountsForProductService(list[i].ShopID, list[i].ID, primitive.NilObjectID, nil)
		ApplyDiscountsToProduct(&list[i], discounts)
	}
	return list, nil
}

// UpdateProductVariantIDs updates existing products to ensure all variants have proper IDs
// This is a one-time migration function to fix existing data
func UpdateProductVariantIDs() error {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return err
	}

	for _, product := range products {
		needsUpdate := false
		updatedVariants := make([]models.Variant, len(product.Variants))

		for i, variant := range product.Variants {
			updatedVariants[i] = variant
			// If variant has no ID but has real options, generate one
			if variant.VariantID.IsZero() && len(variant.Options) > 0 && (variant.Options[0].Name != "" || variant.Options[0].Value != "") {
				updatedVariants[i].VariantID = primitive.NewObjectID()
				needsUpdate = true
			}
		}

		if needsUpdate {
			_, err := repositories.UpdateProduct(product.ID.Hex(), bson.M{
				"variants": updatedVariants,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// EnsureProductVariantIDs ensures that a specific product has proper variant IDs
// This function is called during product retrieval to fix variant IDs on-the-fly
func EnsureProductVariantIDs(product *models.Product) error {
	needsUpdate := false
	updatedVariants := make([]models.Variant, len(product.Variants))

	for i, variant := range product.Variants {
		updatedVariants[i] = variant
		// If variant has no ID but has real options, generate one
		if variant.VariantID.IsZero() && len(variant.Options) > 0 && (variant.Options[0].Name != "" || variant.Options[0].Value != "") {
			newID := primitive.NewObjectID()
			updatedVariants[i].VariantID = newID
			needsUpdate = true
			fmt.Printf("Generated new variant ID %s for product %s variant %d\n", newID.Hex(), product.ID.Hex(), i)
		}
	}

	if needsUpdate {
		_, err := repositories.UpdateProduct(product.ID.Hex(), bson.M{
			"variants": updatedVariants,
		})
		if err != nil {
			return err
		}
		// Update the product in memory with the new variant IDs
		product.Variants = updatedVariants
		fmt.Printf("Updated product %s with %d variants in database\n", product.ID.Hex(), len(updatedVariants))
	}

	return nil
}

// DiscountToAPIResponse converts a discount model to API response format for storefront
func DiscountToAPIResponse(d *models.Discount) map[string]interface{} {
	return map[string]interface{}{
		"id":          d.ID.Hex(),
		"name":        d.Name,
		"description": d.Description,
		"category":    d.Category,
		"type":        d.Type,
		"value":       d.Value,
		"start_at":    d.StartAt,
		"end_at":      d.EndAt,
		"active":      d.Active,
	}
}

// GetActiveDiscountsForProductAPI fetches active discounts for a product and converts them to API format
func GetActiveDiscountsForProductAPI(shopID, productID primitive.ObjectID, variantIDs []primitive.ObjectID, collectionIDs []primitive.ObjectID) ([]map[string]interface{}, error) {
	// Get all active discounts for this product
	discounts, err := GetActiveDiscountsForProductService(shopID, productID, primitive.NilObjectID, collectionIDs)
	if err != nil {
		return nil, err
	}

	// Also get discounts for specific variants
	for _, variantID := range variantIDs {
		variantDiscounts, err := GetActiveDiscountsForProductService(shopID, productID, variantID, collectionIDs)
		if err != nil {
			continue // Skip if error, but continue with other variants
		}
		discounts = append(discounts, variantDiscounts...)
	}

	// Convert to API format and remove duplicates
	seen := make(map[string]bool)
	var apiDiscounts []map[string]interface{}
	for _, d := range discounts {
		if !seen[d.ID.Hex()] {
			seen[d.ID.Hex()] = true
			apiDiscounts = append(apiDiscounts, DiscountToAPIResponse(&d))
		}
	}

	return apiDiscounts, nil
}

func ProductToAPIResponse(p *models.Product) map[string]interface{} {
	resp := map[string]interface{}{
		"id":          p.ID.Hex(),
		"name":        p.Name,
		"slug":        p.Slug,
		"description": p.Description,
		"main_image":  p.MainImage,
		"images":      p.Images,
		"category":    p.Category,
		"createdBy":   p.CreatedBy.Hex(),
		"createdAt":   p.CreatedAt,
		"updatedAt":   p.UpdatedAt,
	}

	// Check if product has real variants (not just empty ones)
	hasRealVariants := false
	if len(p.Variants) > 0 {
		for _, v := range p.Variants {
			// A real variant has meaningful options (not empty name/value)
			if len(v.Options) > 0 && (v.Options[0].Name != "" || v.Options[0].Value != "") {
				hasRealVariants = true
				break
			}
		}
	}

	if hasRealVariants {
		// Product has real variants - include variant information
		realVariants := []models.Variant{}
		for _, v := range p.Variants {
			realVariants = append(realVariants, v)
		}

		// Compute starting_price, highest_display_price, total_stock
		minPrice := realVariants[0].Price
		maxDisplayPrice := realVariants[0].DisplayPrice
		totalStock := 0
		for _, v := range realVariants {
			if v.Price < minPrice {
				minPrice = v.Price
			}
			if v.DisplayPrice != nil && (maxDisplayPrice == nil || *v.DisplayPrice > *maxDisplayPrice) {
				maxDisplayPrice = v.DisplayPrice
			}
			totalStock += v.Stock
		}
		resp["starting_price"] = minPrice
		if maxDisplayPrice != nil {
			resp["highest_display_price"] = *maxDisplayPrice
		}
		resp["total_stock"] = totalStock
		resp["variants"] = realVariants
		return resp
	}

	// No real variants: treat as simple product (no variants array)
	resp["price"] = p.Price
	if p.DisplayPrice != nil {
		resp["display_price"] = *p.DisplayPrice
	}
	resp["stock"] = p.Stock
	return resp
}

// GetCollectionIDsForProduct returns the collection IDs that contain this product
func GetCollectionIDsForProduct(productID primitive.ObjectID) ([]primitive.ObjectID, error) {
	// Use the seller collection repository
	collections, err := sellerRepos.GetCollectionsByFilter(bson.M{"product_ids": productID})
	if err != nil {
		return nil, err
	}

	var collectionIDs []primitive.ObjectID
	for _, coll := range collections {
		collectionIDs = append(collectionIDs, coll.ID)
	}

	return collectionIDs, nil
}

// ProductToAPIResponseWithDiscounts converts a product to API response format and includes active discounts
func ProductToAPIResponseWithDiscounts(p *models.Product) map[string]interface{} {
	// Get base product response
	resp := ProductToAPIResponse(p)

	// Get collection IDs for this product
	collectionIDs, err := GetCollectionIDsForProduct(p.ID)
	if err != nil {
		// If we can't get collection IDs, just use empty slice
		collectionIDs = []primitive.ObjectID{}
	}

	// Get variant IDs for discount lookup
	var variantIDs []primitive.ObjectID
	for _, v := range p.Variants {
		if !v.VariantID.IsZero() {
			variantIDs = append(variantIDs, v.VariantID)
		}
	}

	// Fetch active discounts for this product
	discounts, err := GetActiveDiscountsForProductAPI(p.ShopID, p.ID, variantIDs, collectionIDs)
	if err == nil && len(discounts) > 0 {
		resp["discounts"] = discounts
	}

	return resp
}
