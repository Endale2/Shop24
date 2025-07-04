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

	// 1) If no variants, inject a "default" variant
	if len(p.Variants) == 0 {
		defaultOpt := models.Option{Name: "", Value: ""}
		defaultVar := models.Variant{
			VariantID:    primitive.NewObjectID(),
			Options:      []models.Option{defaultOpt},
			Price:        p.Price,        // use product price as fallback
			DisplayPrice: p.DisplayPrice, // inherit if set
			Stock:        p.Stock,
			Image:        "",
			Total:        nil, // or set to a calculated value
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		p.Variants = []models.Variant{defaultVar}
	}

	// 2) Ensure every variant has IDs and timestamps
	for i := range p.Variants {
		v := &p.Variants[i]
		if v.VariantID.IsZero() {
			v.VariantID = primitive.NewObjectID()
		}
		if v.CreatedAt.IsZero() {
			v.CreatedAt = now
		}
		v.UpdatedAt = now
	}

	if len(p.Variants) > 0 {
		// 3) If product has variants, set price to lowest variant price, display_price to highest variant display_price, and stock to sum of variant stocks. Lock these fields from being edited directly.
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
	} else {
		// 4) If no variants, allow product.price, display_price, and stock to be set/edited directly.
		// (No action needed; use as provided)
	}

	// 5) Determine MainImage
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
		normalizeProduct(p)
	}
	return p, nil
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

// Add a function to produce the correct JSON response for a product
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
	if len(p.Variants) > 0 {
		// Only include variants if they are real (not a default injected one)
		realVariants := []models.Variant{}
		for _, v := range p.Variants {
			if len(v.Options) > 0 && (v.Options[0].Name != "" || v.Options[0].Value != "") {
				realVariants = append(realVariants, v)
			}
		}
		if len(realVariants) > 0 {
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
	}
	// No real variants: treat as plain product
	resp["price"] = p.Price
	if p.DisplayPrice != nil {
		resp["display_price"] = *p.DisplayPrice
	}
	resp["stock"] = p.Stock
	return resp
}
