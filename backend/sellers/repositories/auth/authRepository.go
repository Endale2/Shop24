package auth

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/DPRS/models"
)

// AuthRepository defines the interface for authentication-related data operations.
type AuthRepository interface {
	CreateSeller(ctx context.Context, seller *models.AuthSeller) error
	FindSellerByEmail(ctx context.Context, email string) (*models.AuthSeller, error)
	UpdateLastLogin(ctx context.Context, id primitive.ObjectID, lastLogin time.Time) error
}

type authRepository struct {
	collection *mongo.Collection
}

// NewAuthRepository returns a new instance of AuthRepository.
// It expects a mongo.Database from which it selects the "auth_sellers" collection.
func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &authRepository{
		collection: db.Collection("auth_sellers"),
	}
}

// CreateSeller inserts a new seller into the database.
func (r *authRepository) CreateSeller(ctx context.Context, seller *models.AuthSeller) error {
	seller.ID = primitive.NewObjectID()
	seller.CreatedAt = time.Now()
	seller.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, seller)
	return err
}

// FindSellerByEmail retrieves a seller by email.
func (r *authRepository) FindSellerByEmail(ctx context.Context, email string) (*models.AuthSeller, error) {
	var seller models.AuthSeller
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&seller)
	if err != nil {
		return nil, err
	}
	return &seller, nil
}

// UpdateLastLogin updates the last login time for a seller.
func (r *authRepository) UpdateLastLogin(ctx context.Context, id primitive.ObjectID, lastLogin time.Time) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"last_login": lastLogin,
			"updated_at": time.Now(),
		},
	}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no seller found to update")
	}
	return nil
}
