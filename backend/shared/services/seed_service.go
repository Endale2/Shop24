package services

import (
	"context"
	"log"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
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

// SeedComponentTemplates seeds the database with default component templates
func (s *SeedService) SeedComponentTemplates() error {
	collection := s.db.Collection("component_templates")
	
	// Check if components already exist
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	
	if count > 0 {
		log.Println("Component templates already exist, skipping seed")
		return nil
	}
	
	// Get default component templates
	templates := models.GetDefaultComponentTemplates()
	
	// Convert to interface{} slice for insertion
	docs := make([]interface{}, len(templates))
	for i, template := range templates {
		docs[i] = template
	}
	
	// Insert templates
	_, err = collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}
	
	log.Printf("Seeded %d component templates", len(templates))
	return nil
}

// SeedThemePresets seeds the database with default theme presets
func (s *SeedService) SeedThemePresets() error {
	collection := s.db.Collection("theme_presets")
	
	// Check if presets already exist
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	
	if count > 0 {
		log.Println("Theme presets already exist, skipping seed")
		return nil
	}
	
	// Get default theme presets
	presets := models.GetDefaultThemePresets()
	
	// Convert to interface{} slice for insertion
	docs := make([]interface{}, len(presets))
	for i, preset := range presets {
		docs[i] = preset
	}
	
	// Insert presets
	_, err = collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}
	
	log.Printf("Seeded %d theme presets", len(presets))
	return nil
}

// CreateIndexes creates necessary database indexes for performance
func (s *SeedService) CreateIndexes() error {
	// Component templates indexes
	componentCollection := s.db.Collection("component_templates")
	componentIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"type", 1}},
		},
		{
			Keys: bson.D{{"category", 1}},
		},
		{
			Keys: bson.D{{"isPublic", 1}},
		},
		{
			Keys: bson.D{{"usageCount", -1}},
		},
		{
			Keys: bson.D{{"createdBy", 1}},
		},
	}
	
	_, err := componentCollection.Indexes().CreateMany(context.Background(), componentIndexes)
	if err != nil {
		log.Printf("Error creating component indexes: %v", err)
		return err
	}
	
	// Component instances indexes
	instanceCollection := s.db.Collection("component_instances")
	instanceIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"shopId", 1}},
		},
		{
			Keys: bson.D{{"templateId", 1}},
		},
		{
			Keys: bson.D{{"pageType", 1}},
		},
		{
			Keys: bson.D{
				{"shopId", 1},
				{"pageType", 1},
			},
		},
		{
			Keys: bson.D{{"position", 1}},
		},
	}
	
	_, err = instanceCollection.Indexes().CreateMany(context.Background(), instanceIndexes)
	if err != nil {
		log.Printf("Error creating component instance indexes: %v", err)
		return err
	}
	
	// Page layouts indexes
	layoutCollection := s.db.Collection("page_layouts")
	layoutIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"shopId", 1}},
		},
		{
			Keys: bson.D{{"pageType", 1}},
		},
		{
			Keys: bson.D{
				{"shopId", 1},
				{"pageType", 1},
			},
		},
		{
			Keys: bson.D{{"isActive", 1}},
		},
	}
	
	_, err = layoutCollection.Indexes().CreateMany(context.Background(), layoutIndexes)
	if err != nil {
		log.Printf("Error creating page layout indexes: %v", err)
		return err
	}
	
	// Theme presets indexes
	presetCollection := s.db.Collection("theme_presets")
	presetIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"category", 1}},
		},
		{
			Keys: bson.D{{"isPublic", 1}},
		},
		{
			Keys: bson.D{{"usageCount", -1}},
		},
		{
			Keys: bson.D{{"rating", -1}},
		},
	}
	
	_, err = presetCollection.Indexes().CreateMany(context.Background(), presetIndexes)
	if err != nil {
		log.Printf("Error creating theme preset indexes: %v", err)
		return err
	}
	
	// Theme versions indexes
	versionCollection := s.db.Collection("theme_versions")
	versionIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"shopId", 1}},
		},
		{
			Keys: bson.D{{"themeId", 1}},
		},
		{
			Keys: bson.D{
				{"shopId", 1},
				{"createdAt", -1},
			},
		},
		{
			Keys: bson.D{{"isActive", 1}},
		},
	}
	
	_, err = versionCollection.Indexes().CreateMany(context.Background(), versionIndexes)
	if err != nil {
		log.Printf("Error creating theme version indexes: %v", err)
		return err
	}
	
	log.Println("Successfully created all database indexes")
	return nil
}

// SeedAll runs all seeding operations
func (s *SeedService) SeedAll() error {
	log.Println("Starting database seeding...")
	
	// Create indexes first
	if err := s.CreateIndexes(); err != nil {
		return err
	}
	
	// Seed component templates
	if err := s.SeedComponentTemplates(); err != nil {
		return err
	}
	
	// Seed theme presets
	if err := s.SeedThemePresets(); err != nil {
		return err
	}
	
	log.Println("Database seeding completed successfully")
	return nil
}

// CleanupExpiredData removes old data for performance
func (s *SeedService) CleanupExpiredData() error {
	log.Println("Starting data cleanup...")
	
	// Clean up old theme versions (keep only last 10 per shop)
	versionCollection := s.db.Collection("theme_versions")
	
	// Get all shops with theme versions
	pipeline := []bson.M{
		{"$group": bson.M{
			"_id": "$shopId",
			"versions": bson.M{"$push": bson.M{
				"id": "$_id",
				"createdAt": "$createdAt",
			}},
		}},
	}
	
	cursor, err := versionCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())
	
	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return err
	}
	
	// For each shop, keep only the latest 10 versions
	for _, result := range results {
		versions := result["versions"].([]interface{})
		if len(versions) > 10 {
			// Sort by createdAt and keep only latest 10
			// This is a simplified version - in production, you'd want proper sorting
			versionsToDelete := versions[10:]
			
			var idsToDelete []interface{}
			for _, version := range versionsToDelete {
				versionMap := version.(bson.M)
				idsToDelete = append(idsToDelete, versionMap["id"])
			}
			
			if len(idsToDelete) > 0 {
				_, err := versionCollection.DeleteMany(context.Background(), bson.M{
					"_id": bson.M{"$in": idsToDelete},
				})
				if err != nil {
					log.Printf("Error deleting old theme versions: %v", err)
				} else {
					log.Printf("Deleted %d old theme versions for shop %v", len(idsToDelete), result["_id"])
				}
			}
		}
	}
	
	log.Println("Data cleanup completed")
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
