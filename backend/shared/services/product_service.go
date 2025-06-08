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
	s = strings.Join(strings.Fields(s), "-")
	return s
}

// normalizeProduct ensures that:
//  1) If no variants exist, we inject a “default” variant
//  2) Each variant has a non‐zero VariantID, CreatedAt, UpdatedAt
//  3) The Product’s DisplayPrice is set to the minimum Variant price
//  4) The Product’s MainImage is set from Images[0] or first variant image
func normalizeProduct(p *models.Product) {
	now := time.Now()

	// 1) If no variants, inject a “default” variant
	if len(p.Variants) == 0 {
		defaultVar := models.Variant{
			VariantID: primitive.NewObjectID(),
			Options:   map[string]string{},
			Price:     p.DisplayPrice,
			Stock:     0,
			Image:     "", // will be filled below
			CreatedAt: now,
			UpdatedAt: now,
		}
		p.Variants = []models.Variant{defaultVar}
	}

	// 2) Ensure every variant has IDs and timestamps
	for i := range p.Variants {
		if p.Variants[i].VariantID.IsZero() {
			p.Variants[i].VariantID = primitive.NewObjectID()
		}
		if p.Variants[i].CreatedAt.IsZero() {
			p.Variants[i].CreatedAt = now
		}
		p.Variants[i].UpdatedAt = now
	}

	// 3) Compute DisplayPrice = minimum variant price
	minPrice := p.Variants[0].Price
	for _, v := range p.Variants[1:] {
		if v.Price < minPrice {
			minPrice = v.Price
		}
	}
	p.DisplayPrice = minPrice

	// 4) Determine MainImage
	if len(p.Images) > 0 {
		p.MainImage = p.Images[0]
	} else {
		// fallback: first non‐empty variant image
		for _, v := range p.Variants {
			if v.Image != "" {
				p.MainImage = v.Image
				break
			}
		}
	}
}

// ------------------------------------------------------------
// Service: Create, Read, Update, Delete, plus “by slug”
// ------------------------------------------------------------

// CreateProductService inserts a new Product into MongoDB after applying
// business logic (ID, slug, normalization including MainImage).
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

	// Stamp variant IDs & timestamps if provided
	for i := range p.Variants {
		if p.Variants[i].VariantID.IsZero() {
			p.Variants[i].VariantID = primitive.NewObjectID()
		}
		if p.Variants[i].CreatedAt.IsZero() {
			p.Variants[i].CreatedAt = now
		}
		p.Variants[i].UpdatedAt = now
	}

	if p.DisplayPrice == 0 {
		p.DisplayPrice = p.Price
	}

	// Normalize (sets DisplayPrice, ensures variants, sets MainImage)
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
