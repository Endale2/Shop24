package services

import (
	"github.com/Endale2/DRPS/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// SeedService handles database seeding operations
type SeedService struct {
	db *mongo.Database
}

// NewSeedService creates a new seed service instance
func NewSeedService() *SeedService {
	return &SeedService{
		db: config.GetDB(),
	}
}

// CreateIndexes creates necessary database indexes for performance
// Note: Theme/customization-related indexes removed.
func (s *SeedService) CreateIndexes() error {
	return nil
}

// SeedAll runs all seeding operations
// Note: Theme/customization-related seeders removed.
func (s *SeedService) SeedAll() error {
	return nil
}

// CleanupExpiredData removes old data for performance
// Note: Theme version cleanup removed.
func (s *SeedService) CleanupExpiredData() error {
	return nil
}

// Global seed service instance
var seedService *SeedService

// GetSeedService returns the global seed service instance
func GetSeedService() *SeedService {
	if seedService == nil {
		seedService = NewSeedService()
	}
	return seedService
}
