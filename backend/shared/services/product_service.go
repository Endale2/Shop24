package services

import (
	"github.com/Endale2/DRPS/shared/models"
	"github.com/Endale2/DRPS/shared/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateProductService creates a new product.
func CreateProductService(product *models.Product) (*mongo.InsertOneResult, error) {
	// You can add business logic or validations here if needed.
	return repositories.CreateProduct(product)
}

// GetProductByIDService retrieves a product by its ID.
func GetProductByIDService(id string) (*models.Product, error) {
	return repositories.GetProductByID(id)
}

// GetAllProductsService returns all products.
func GetAllProductsService() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

// UpdateProductService updates fields of a product identified by its ID.
func UpdateProductService(id string, updatedData bson.M) (*mongo.UpdateResult, error) {
	// Add any service-level validations or business rules here.
	return repositories.UpdateProduct(id, updatedData)
}

// DeleteProductService removes a product by its ID.
func DeleteProductService(id string) (*mongo.DeleteResult, error) {
	return repositories.DeleteProduct(id)
}



// CountProductsService returns the total number of products matching the given filter.
func CountProductsService(filter bson.M) (int64, error) {
    return repositories.CountProducts(filter)
}

// CountByCategoryService returns counts of products grouped by category.
func CountByCategoryService() (map[string]int64, error) {
    return repositories.CountByCategory()
}


// GetProductsByShopIDService retrieves all products for a specific shop.
func GetProductsByShopIDService(shopID string) ([]models.Product, error) {
	// 1) convert hex string to ObjectID
	oid, err := primitive.ObjectIDFromHex(shopID)
	if err != nil {
		return nil, err
	}
	// 2) filter on the exact BSON field name: "shop_id"
	filter := bson.M{"shop_id": oid}
	// 3) delegate to the repository
	return repositories.GetProductsByFilter(filter)
}