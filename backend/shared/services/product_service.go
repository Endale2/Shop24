package services

import (
	"errors"
	"strings"
	"time"

	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ------------------------------------------------------------
// Helpers: slugify, normalize, etc.
// ------------------------------------------------------------

// slugify takes a name (e.g. "Catan Base Game") and returns
// a simple URL‐friendly slug ("catan-base-game").
func slugify(name string) string {
	s := strings.ToLower(name)
	// Replace spaces (and multiple spaces) with a single hyphen
	s = strings.Join(strings.Fields(s), "-")
	// Optionally: strip out any unsupported characters (e.g. punctuation).
	// For now, we’ll assume Fields() has removed most whitespace issues.
	return s
}

// normalizeProduct ensures that:
//  1) If no variants exist, we inject a “default” variant
//  2) Each variant has a non‐zero VariantID, CreatedAt, UpdatedAt
//  3) The Product’s DisplayPrice is set to the minimum Variant price
func normalizeProduct(p *models.Product) {
	now := time.Now()

	// 1) If no variants, inject a “default” variant that
	//    slides in with the top‐level Price (or DisplayPrice if set).
	if len(p.Variants) == 0 {
		defaultVar := models.Variant{
			VariantID: primitive.NewObjectID(),
			Options:   map[string]string{},
			Price:     p.DisplayPrice,
			Stock:     0,
			Image:     p.MainImage,
			CreatedAt: now,
			UpdatedAt: now,
		}
		p.Variants = []models.Variant{defaultVar}
	}

	// 2) Ensure every variant has a valid VariantID and timestamps.
	for i := range p.Variants {
		if p.Variants[i].VariantID.IsZero() {
			p.Variants[i].VariantID = primitive.NewObjectID()
		}
		// If you store CreatedAt/UpdatedAt on Variant, set them here if zero.
		if p.Variants[i].CreatedAt.IsZero() {
			p.Variants[i].CreatedAt = now
		}
		p.Variants[i].UpdatedAt = now
	}

	// 3) Compute DisplayPrice = minimum of all variant.Price
	minPrice := p.Variants[0].Price
	for _, v := range p.Variants[1:] {
		if v.Price < minPrice {
			minPrice = v.Price
		}
	}
	p.DisplayPrice = minPrice
}

// ------------------------------------------------------------
// Service: Create, Read, Update, Delete, plus “by slug”
// ------------------------------------------------------------

// CreateProductService inserts a new Product into MongoDB after applying
// all necessary business logic (ID assignment, slug generation, normalization).
func CreateProductService(p *models.Product) (*mongo.InsertOneResult, error) {
	// 1) Basic validation
	if strings.TrimSpace(p.Name) == "" {
		return nil, errors.New("product name is required")
	}

	now := time.Now()

	// 2) Assign a new ObjectID to the top‐level product
	p.ID = primitive.NewObjectID()
	p.CreatedAt = now
	p.UpdatedAt = now

	// 3) If no slug was provided, generate one from the name.
	if strings.TrimSpace(p.Slug) == "" {
		p.Slug = slugify(p.Name)
	}

	// 4) Assign each Variant a new VariantID if zero, and stamp timestamps.
	for i := range p.Variants {
		if p.Variants[i].VariantID.IsZero() {
			p.Variants[i].VariantID = primitive.NewObjectID()
		}
		// Set variant timestamps if not already set
		if p.Variants[i].CreatedAt.IsZero() {
			p.Variants[i].CreatedAt = now
		}
		p.Variants[i].UpdatedAt = now
	}

	// 5) If top‐level DisplayPrice is zero, set it from p.Price
	if p.DisplayPrice == 0 {
		p.DisplayPrice = p.Price
	}

	// 6) Normalize variants & DisplayPrice
	normalizeProduct(p)

	// 7) Delegate to repository
	return repositories.CreateProduct(p)
}

// GetProductByIDService retrieves a Product by its hex ID, then normalizes it.
func GetProductByIDService(id string) (*models.Product, error) {
	p, err := repositories.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if p != nil {
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

// UpdateProductService updates fields by ID, then (optionally) you can choose
// to reload & normalize if critical fields changed. For now, it just updates.
func UpdateProductService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	updatedData["updatedAt"] = time.Now()

	// If the client attempts to change the Name, you may want to regenerate slug:
	if newName, ok := updatedData["name"].(string); ok && strings.TrimSpace(newName) != "" {
		updatedData["slug"] = slugify(newName)
	}

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
	// repositories.GetProductsByShopSlug already filters by shop_slug → shop_id
	for i := range list {
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
	if p == nil {
		return nil, nil
	}
	normalizeProduct(p)
	return p, nil
}
