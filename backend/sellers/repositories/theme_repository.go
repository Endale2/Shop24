package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Endale2/DRPS/config"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var themeCollection *mongo.Collection = config.GetCollection("DRPS", "themes")
var shopThemeCollection *mongo.Collection = config.GetCollection("DRPS", "shop_themes")
var themePresetCollection *mongo.Collection = config.GetCollection("DRPS", "theme_presets")

// =============== Theme Base Operations ===============

// CreateTheme inserts a new Theme document
func CreateTheme(theme *models.Theme) (*mongo.InsertOneResult, error) {
	theme.ID = primitive.NewObjectID()
	theme.CreatedAt = time.Now()
	theme.UpdatedAt = time.Now()

	return themeCollection.InsertOne(context.Background(), theme)
}

// GetThemeByID retrieves a Theme by its ObjectID
func GetThemeByID(id primitive.ObjectID) (*models.Theme, error) {
	var theme models.Theme
	err := themeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&theme)
	if err != nil {
		return nil, err
	}
	return &theme, nil
}

// GetPublicThemes returns all public/marketplace themes
func GetPublicThemes() ([]models.Theme, error) {
	cursor, err := themeCollection.Find(context.Background(), bson.M{"isPublic": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.Theme
	for cursor.Next(context.Background()) {
		var theme models.Theme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

// GetThemesByCreator returns themes created by a specific seller
func GetThemesByCreator(creatorID primitive.ObjectID) ([]models.Theme, error) {
	cursor, err := themeCollection.Find(context.Background(), bson.M{"createdBy": creatorID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.Theme
	for cursor.Next(context.Background()) {
		var theme models.Theme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

// UpdateTheme updates a theme document
func UpdateTheme(id primitive.ObjectID, updatedData bson.M) (*mongo.UpdateResult, error) {
	if id.IsZero() {
		return nil, errors.New("invalid theme ID")
	}
	updatedData["updatedAt"] = time.Now()

	return themeCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
}

// DeleteTheme removes a theme document
func DeleteTheme(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return themeCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// =============== Shop Theme Operations ===============

// CreateShopTheme creates a new ShopTheme configuration
func CreateShopTheme(shopTheme *models.ShopTheme) (*mongo.InsertOneResult, error) {
	shopTheme.ID = primitive.NewObjectID()
	shopTheme.CreatedAt = time.Now()
	shopTheme.UpdatedAt = time.Now()

	return shopThemeCollection.InsertOne(context.Background(), shopTheme)
}

// GetShopThemeByShopID retrieves the active theme for a shop
func GetShopThemeByShopID(shopID primitive.ObjectID) (*models.ShopTheme, error) {
	var shopTheme models.ShopTheme
	err := shopThemeCollection.FindOne(context.Background(), bson.M{
		"shopId":    shopID,
		"published": true,
	}).Decode(&shopTheme)
	if err != nil {
		return nil, err
	}
	return &shopTheme, nil
}

// GetShopThemeByID retrieves a specific ShopTheme by ID
func GetShopThemeByID(id primitive.ObjectID) (*models.ShopTheme, error) {
	var shopTheme models.ShopTheme
	err := shopThemeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&shopTheme)
	if err != nil {
		return nil, err
	}
	return &shopTheme, nil
}

// GetShopThemesByShopID retrieves all theme configurations for a shop (published and unpublished)
func GetShopThemesByShopID(shopID primitive.ObjectID) ([]models.ShopTheme, error) {
	cursor, err := shopThemeCollection.Find(context.Background(), bson.M{"shopId": shopID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var shopThemes []models.ShopTheme
	for cursor.Next(context.Background()) {
		var shopTheme models.ShopTheme
		if err := cursor.Decode(&shopTheme); err != nil {
			return nil, err
		}
		shopThemes = append(shopThemes, shopTheme)
	}
	return shopThemes, nil
}

// UpdateShopTheme updates a ShopTheme configuration
func UpdateShopTheme(id primitive.ObjectID, updatedData bson.M) (*mongo.UpdateResult, error) {
	if id.IsZero() {
		return nil, errors.New("invalid shop theme ID")
	}
	updatedData["updatedAt"] = time.Now()

	return shopThemeCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
}

// PublishShopTheme sets a theme as published and unpublishes others for the same shop
func PublishShopTheme(shopID, shopThemeID primitive.ObjectID) error {
	// Start a transaction to ensure atomicity
	session, err := config.DB.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		// First, unpublish all themes for this shop
		_, err := shopThemeCollection.UpdateMany(
			sessionContext,
			bson.M{"shopId": shopID},
			bson.M{"$set": bson.M{"published": false, "updatedAt": time.Now()}},
		)
		if err != nil {
			return nil, err
		}

		// Then publish the selected theme
		_, err = shopThemeCollection.UpdateOne(
			sessionContext,
			bson.M{"_id": shopThemeID, "shopId": shopID},
			bson.M{"$set": bson.M{"published": true, "updatedAt": time.Now()}},
		)
		return nil, err
	}

	_, err = session.WithTransaction(context.Background(), callback)
	return err
}

// DeleteShopTheme removes a ShopTheme configuration
func DeleteShopTheme(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return shopThemeCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// =============== Theme Preset Operations ===============

// CreateThemePreset creates a new theme preset
func CreateThemePreset(preset *models.ThemePreset) (*mongo.InsertOneResult, error) {
	preset.ID = primitive.NewObjectID()
	preset.CreatedAt = time.Now()
	preset.UpdatedAt = time.Now()

	return themePresetCollection.InsertOne(context.Background(), preset)
}

// GetThemePresetsByThemeID retrieves all presets for a specific theme
func GetThemePresetsByThemeID(themeID primitive.ObjectID) ([]models.ThemePreset, error) {
	cursor, err := themePresetCollection.Find(context.Background(), bson.M{"themeId": themeID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var presets []models.ThemePreset
	for cursor.Next(context.Background()) {
		var preset models.ThemePreset
		if err := cursor.Decode(&preset); err != nil {
			return nil, err
		}
		presets = append(presets, preset)
	}
	return presets, nil
}

// GetThemePresetByID retrieves a specific preset by ID
func GetThemePresetByID(id primitive.ObjectID) (*models.ThemePreset, error) {
	var preset models.ThemePreset
	err := themePresetCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&preset)
	if err != nil {
		return nil, err
	}
	return &preset, nil
}

// GetDefaultThemePreset retrieves the default preset for a theme
func GetDefaultThemePreset(themeID primitive.ObjectID) (*models.ThemePreset, error) {
	var preset models.ThemePreset
	err := themePresetCollection.FindOne(context.Background(), bson.M{
		"themeId":   themeID,
		"isDefault": true,
	}).Decode(&preset)
	if err != nil {
		return nil, err
	}
	return &preset, nil
}

// UpdateThemePreset updates a theme preset
func UpdateThemePreset(id primitive.ObjectID, updatedData bson.M) (*mongo.UpdateResult, error) {
	if id.IsZero() {
		return nil, errors.New("invalid preset ID")
	}
	updatedData["updatedAt"] = time.Now()

	return themePresetCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
}

// DeleteThemePreset removes a theme preset
func DeleteThemePreset(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return themePresetCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// =============== Search and Filter Operations ===============

// SearchThemes searches themes by name, tags, or description
func SearchThemes(query string, isPublic bool) ([]models.Theme, error) {
	filter := bson.M{
		"isPublic": isPublic,
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
			{"tags": bson.M{"$in": []string{query}}},
		},
	}

	cursor, err := themeCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.Theme
	for cursor.Next(context.Background()) {
		var theme models.Theme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

// GetThemesByTags retrieves themes filtered by tags
func GetThemesByTags(tags []string, isPublic bool) ([]models.Theme, error) {
	filter := bson.M{
		"isPublic": isPublic,
		"tags":     bson.M{"$in": tags},
	}

	cursor, err := themeCollection.Find(context.Background(), filter, options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.Theme
	for cursor.Next(context.Background()) {
		var theme models.Theme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}
