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

// ShopTheme collection for scalable theme management
var shopThemeCollection *mongo.Collection = config.GetCollection("DRPS", "shop_themes")

// =============== CRUD Operations ===============

// CreateShopTheme creates a new shop theme
func CreateShopTheme(theme *models.ShopTheme) (*mongo.InsertOneResult, error) {
	if theme.ID.IsZero() {
		theme.ID = primitive.NewObjectID()
	}
	theme.CreatedAt = time.Now()
	theme.UpdatedAt = time.Now()

	return shopThemeCollection.InsertOne(context.Background(), theme)
}

// GetShopThemeByID retrieves a shop theme by its ID
func GetShopThemeByID(id primitive.ObjectID) (*models.ShopTheme, error) {
	var theme models.ShopTheme
	err := shopThemeCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&theme)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &theme, nil
}

// GetActiveShopTheme retrieves the active theme for a shop
func GetActiveShopTheme(shopID primitive.ObjectID) (*models.ShopTheme, error) {
	var theme models.ShopTheme
	filter := bson.M{
		"shopId":   shopID,
		"isActive": true,
	}
	
	err := shopThemeCollection.FindOne(context.Background(), filter).Decode(&theme)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &theme, nil
}

// GetShopThemes retrieves all themes for a shop
func GetShopThemes(shopID primitive.ObjectID) ([]models.ShopTheme, error) {
	filter := bson.M{"shopId": shopID}
	opts := options.Find().SetSort(bson.D{{"createdAt", -1}}) // Latest first

	cursor, err := shopThemeCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.ShopTheme
	for cursor.Next(context.Background()) {
		var theme models.ShopTheme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

// UpdateShopTheme updates an existing shop theme
func UpdateShopTheme(id primitive.ObjectID, updateData bson.M) (*mongo.UpdateResult, error) {
	updateData["updatedAt"] = time.Now()
	
	return shopThemeCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updateData},
	)
}

// DeleteShopTheme deletes a shop theme
func DeleteShopTheme(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return shopThemeCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// =============== Advanced Operations ===============

// SetActiveTheme sets a theme as active and deactivates others for the same shop
func SetActiveTheme(shopID, themeID primitive.ObjectID) error {
	// Start a transaction to ensure atomicity
	session, err := config.DB.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		// First, deactivate all themes for this shop
		_, err := shopThemeCollection.UpdateMany(
			sessionContext,
			bson.M{"shopId": shopID},
			bson.M{"$set": bson.M{"isActive": false, "updatedAt": time.Now()}},
		)
		if err != nil {
			return nil, err
		}

		// Then activate the selected theme
		_, err = shopThemeCollection.UpdateOne(
			sessionContext,
			bson.M{"_id": themeID, "shopId": shopID},
			bson.M{"$set": bson.M{
				"isActive":   true,
				"lastUsedAt": time.Now(),
				"updatedAt":  time.Now(),
				"$inc":       bson.M{"usageCount": 1},
			}},
		)
		return nil, err
	}

	_, err = session.WithTransaction(context.Background(), callback)
	return err
}

// UpdateThemeColors updates only the colors section of a theme
func UpdateThemeColors(shopID primitive.ObjectID, colors map[string]string) (*models.ShopTheme, error) {
	filter := bson.M{"shopId": shopID, "isActive": true}
	update := bson.M{
		"$set": bson.M{
			"colors":    colors,
			"updatedAt": time.Now(),
		},
	}

	// Find and update the active theme
	var theme models.ShopTheme
	err := shopThemeCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&theme)
	
	if err != nil {
		return nil, err
	}
	return &theme, nil
}

// UpdateThemeFonts updates only the fonts section of a theme
func UpdateThemeFonts(shopID primitive.ObjectID, fonts map[string]string) (*models.ShopTheme, error) {
	filter := bson.M{"shopId": shopID, "isActive": true}
	update := bson.M{
		"$set": bson.M{
			"fonts":     fonts,
			"updatedAt": time.Now(),
		},
	}

	var theme models.ShopTheme
	err := shopThemeCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&theme)
	
	if err != nil {
		return nil, err
	}
	return &theme, nil
}

// UpdateThemeLayout updates only the layout section of a theme
func UpdateThemeLayout(shopID primitive.ObjectID, layout map[string]string) (*models.ShopTheme, error) {
	filter := bson.M{"shopId": shopID, "isActive": true}
	update := bson.M{
		"$set": bson.M{
			"layout":    layout,
			"updatedAt": time.Now(),
		},
	}

	var theme models.ShopTheme
	err := shopThemeCollection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&theme)
	
	if err != nil {
		return nil, err
	}
	return &theme, nil
}

// CreateBackupAndUpdate creates a backup before updating theme
func CreateBackupAndUpdate(themeID primitive.ObjectID, updateData bson.M) error {
	// Get current theme
	currentTheme, err := GetShopThemeByID(themeID)
	if err != nil {
		return err
	}
	if currentTheme == nil {
		return errors.New("theme not found")
	}

	// Create backup
	currentTheme.CreateBackup()
	
	// Include backup in update
	updateData["backup"] = currentTheme.Backup
	updateData["updatedAt"] = time.Now()

	// Increment version
	currentTheme.IncrementVersion()
	updateData["version"] = currentTheme.Version

	_, err = UpdateShopTheme(themeID, updateData)
	return err
}

// ResetThemeToDefault resets a shop's theme to default configuration
func ResetThemeToDefault(shopID primitive.ObjectID, createdBy primitive.ObjectID) (*models.ShopTheme, error) {
	// Get current active theme for backup
	currentTheme, err := GetActiveShopTheme(shopID)
	if err != nil {
		return nil, err
	}

	var updateData bson.M
	
	if currentTheme != nil {
		// Create backup of current theme
		currentTheme.CreateBackup()
		
		// Reset to default values
		defaultTheme := models.GetDefaultShopTheme(shopID, createdBy)
		updateData = bson.M{
			"colors":    defaultTheme.Colors,
			"fonts":     defaultTheme.Fonts,
			"layout":    defaultTheme.Layout,
			"customCSS": "",
			"seo":       defaultTheme.SEO,
			"mobile":    defaultTheme.Mobile,
			"backup":    currentTheme.Backup,
			"version":   "1.0.0",
			"updatedAt": time.Now(),
		}

		_, err = UpdateShopTheme(currentTheme.ID, updateData)
		if err != nil {
			return nil, err
		}

		// Return updated theme
		return GetShopThemeByID(currentTheme.ID)
	} else {
		// No theme exists, create default one
		defaultTheme := models.GetDefaultShopTheme(shopID, createdBy)
		_, err = CreateShopTheme(defaultTheme)
		if err != nil {
			return nil, err
		}
		return defaultTheme, nil
	}
}

// =============== Analytics & Performance ===============

// UpdateThemeAnalytics updates theme performance metrics
func UpdateThemeAnalytics(themeID primitive.ObjectID, loadTime float64, conversions int, views int) error {
	update := bson.M{
		"$set": bson.M{
			"loadTime":  loadTime,
			"updatedAt": time.Now(),
		},
		"$inc": bson.M{
			"conversions": conversions,
			"views":       views,
		},
	}

	_, err := shopThemeCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": themeID},
		update,
	)
	return err
}

// GetPopularThemes returns most used themes across all shops
func GetPopularThemes(limit int) ([]models.ShopTheme, error) {
	opts := options.Find().
		SetSort(bson.D{{"usageCount", -1}}).
		SetLimit(int64(limit))

	cursor, err := shopThemeCollection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.ShopTheme
	for cursor.Next(context.Background()) {
		var theme models.ShopTheme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

// =============== Cache & Performance ===============

// UpdateThemeCache updates the compiled CSS and cache key
func UpdateThemeCache(themeID primitive.ObjectID, compiledCSS, cacheKey string) error {
	update := bson.M{
		"$set": bson.M{
			"compiledCSS":  compiledCSS,
			"cacheKey":     cacheKey,
			"lastCompiled": time.Now(),
			"updatedAt":    time.Now(),
		},
	}

	_, err := shopThemeCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": themeID},
		update,
	)
	return err
}

// GetThemesByTags retrieves themes by tags
func GetThemesByTags(tags []string, limit int) ([]models.ShopTheme, error) {
	filter := bson.M{"tags": bson.M{"$in": tags}}
	opts := options.Find().SetLimit(int64(limit))

	cursor, err := shopThemeCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var themes []models.ShopTheme
	for cursor.Next(context.Background()) {
		var theme models.ShopTheme
		if err := cursor.Decode(&theme); err != nil {
			return nil, err
		}
		themes = append(themes, theme)
	}
	return themes, nil
}

// =============== Indexing for Performance ===============

// CreateShopThemeIndexes creates necessary indexes for optimal performance
func CreateShopThemeIndexes() error {
	indexes := []mongo.IndexModel{
		// Compound index for shop queries
		{
			Keys: bson.D{
				{"shopId", 1},
				{"isActive", 1},
			},
		},
		// Index for theme lookups
		{
			Keys: bson.D{{"shopId", 1}},
		},
		// Index for active themes
		{
			Keys: bson.D{{"isActive", 1}},
		},
		// Index for analytics
		{
			Keys: bson.D{{"usageCount", -1}},
		},
		// Index for tags
		{
			Keys: bson.D{{"tags", 1}},
		},
		// Index for performance queries
		{
			Keys: bson.D{{"loadTime", 1}},
		},
	}

	_, err := shopThemeCollection.Indexes().CreateMany(context.Background(), indexes)
	return err
}
