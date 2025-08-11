package services

import (
	"errors"
	"strings"

	"github.com/Endale2/DRPS/sellers/repositories"
	"github.com/Endale2/DRPS/shared/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// =============== Theme Base Services ===============

// CreateThemeService creates a new theme with validation
func CreateThemeService(theme *models.Theme, creatorID primitive.ObjectID) (*mongo.InsertOneResult, error) {
	// Validate theme data
	if err := validateTheme(theme); err != nil {
		return nil, err
	}

	// Set creator if creating custom theme
	if !theme.IsPublic {
		theme.CreatedBy = creatorID
	}

	// Set default values
	if theme.Version == "" {
		theme.Version = "1.0.0"
	}

	return repositories.CreateTheme(theme)
}

// GetThemeByIDService retrieves a theme by its ID string
func GetThemeByIDService(idStr string) (*models.Theme, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid theme ID")
	}
	return repositories.GetThemeByID(id)
}

// GetPublicThemesService returns all marketplace themes
func GetPublicThemesService() ([]models.Theme, error) {
	return repositories.GetPublicThemes()
}

// GetThemesByCreatorService returns themes created by a specific seller
func GetThemesByCreatorService(creatorIDStr string) ([]models.Theme, error) {
	creatorID, err := primitive.ObjectIDFromHex(creatorIDStr)
	if err != nil {
		return nil, errors.New("invalid creator ID")
	}
	return repositories.GetThemesByCreator(creatorID)
}

// UpdateThemeService updates a theme with validation
func UpdateThemeService(idStr string, updatedData bson.M, userID primitive.ObjectID) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid theme ID")
	}

	// Check if theme exists and user has permission to update
	theme, err := repositories.GetThemeByID(id)
	if err != nil {
		return nil, errors.New("theme not found")
	}

	// Only allow updates if theme is created by the user or if it's a public theme and user is admin
	if !theme.IsPublic && theme.CreatedBy != userID {
		return nil, errors.New("unauthorized to update this theme")
	}

	// Validate updated fields
	if name, exists := updatedData["name"]; exists {
		if nameStr, ok := name.(string); ok {
			if err := validateThemeName(nameStr); err != nil {
				return nil, err
			}
		}
	}

	return repositories.UpdateTheme(id, updatedData)
}

// DeleteThemeService removes a theme with validation
func DeleteThemeService(idStr string, userID primitive.ObjectID) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, errors.New("invalid theme ID")
	}

	// Check if theme exists and user has permission to delete
	theme, err := repositories.GetThemeByID(id)
	if err != nil {
		return nil, errors.New("theme not found")
	}

	// Only allow deletion if theme is created by the user
	if theme.IsPublic || theme.CreatedBy != userID {
		return nil, errors.New("unauthorized to delete this theme")
	}

	return repositories.DeleteTheme(id)
}

// SearchThemesService searches themes with filters
func SearchThemesService(query string, isPublic bool) ([]models.Theme, error) {
	if strings.TrimSpace(query) == "" {
		if isPublic {
			return repositories.GetPublicThemes()
		}
		return []models.Theme{}, nil
	}
	return repositories.SearchThemes(query, isPublic)
}

// GetThemesByTagsService retrieves themes by tags
func GetThemesByTagsService(tags []string, isPublic bool) ([]models.Theme, error) {
	if len(tags) == 0 {
		return []models.Theme{}, nil
	}
	return repositories.GetThemesByTags(tags, isPublic)
}

// =============== Shop Theme Services ===============

// CreateShopThemeService creates a new shop theme configuration
func CreateShopThemeService(shopTheme *models.ShopTheme, sellerID primitive.ObjectID) (*mongo.InsertOneResult, error) {
	// Validate shop ownership
	if err := validateShopOwnership(shopTheme.ShopID, sellerID); err != nil {
		return nil, err
	}

	// Validate theme exists
	if _, err := repositories.GetThemeByID(shopTheme.ThemeID); err != nil {
		return nil, errors.New("theme not found")
	}

	// Validate theme configuration
	if err := validateThemeConfig(&shopTheme.Overrides); err != nil {
		return nil, err
	}

	return repositories.CreateShopTheme(shopTheme)
}

// GetShopActiveThemeService retrieves the currently active theme for a shop
func GetShopActiveThemeService(shopIDStr string, sellerID primitive.ObjectID) (*models.ShopTheme, *models.Theme, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, nil, errors.New("invalid shop ID")
	}

	// Validate shop ownership
	if err := validateShopOwnership(shopID, sellerID); err != nil {
		return nil, nil, err
	}

	// Get active shop theme
	shopTheme, err := repositories.GetShopThemeByShopID(shopID)
	if err != nil {
		return nil, nil, errors.New("no active theme found for shop")
	}

	// Get base theme details
	theme, err := repositories.GetThemeByID(shopTheme.ThemeID)
	if err != nil {
		return shopTheme, nil, errors.New("base theme not found")
	}

	return shopTheme, theme, nil
}

// GetShopThemesService retrieves all theme configurations for a shop
func GetShopThemesService(shopIDStr string, sellerID primitive.ObjectID) ([]models.ShopTheme, error) {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid shop ID")
	}

	// Validate shop ownership
	if err := validateShopOwnership(shopID, sellerID); err != nil {
		return nil, err
	}

	return repositories.GetShopThemesByShopID(shopID)
}

// UpdateShopThemeService updates a shop theme configuration
func UpdateShopThemeService(shopThemeIDStr string, updatedData bson.M, sellerID primitive.ObjectID) (*mongo.UpdateResult, error) {
	shopThemeID, err := primitive.ObjectIDFromHex(shopThemeIDStr)
	if err != nil {
		return nil, errors.New("invalid shop theme ID")
	}

	// Get existing shop theme
	shopTheme, err := repositories.GetShopThemeByID(shopThemeID)
	if err != nil {
		return nil, errors.New("shop theme not found")
	}

	// Validate shop ownership
	if err := validateShopOwnership(shopTheme.ShopID, sellerID); err != nil {
		return nil, err
	}

	// Validate theme configuration if provided
	if overrides, exists := updatedData["overrides"]; exists {
		if config, ok := overrides.(models.ThemeConfig); ok {
			if err := validateThemeConfig(&config); err != nil {
				return nil, err
			}
		}
	}

	return repositories.UpdateShopTheme(shopThemeID, updatedData)
}

// PublishShopThemeService activates a theme for a shop
func PublishShopThemeService(shopIDStr, shopThemeIDStr string, sellerID primitive.ObjectID) error {
	shopID, err := primitive.ObjectIDFromHex(shopIDStr)
	if err != nil {
		return errors.New("invalid shop ID")
	}

	shopThemeID, err := primitive.ObjectIDFromHex(shopThemeIDStr)
	if err != nil {
		return errors.New("invalid shop theme ID")
	}

	// Validate shop ownership
	if err := validateShopOwnership(shopID, sellerID); err != nil {
		return err
	}

	// Verify shop theme belongs to the shop
	shopTheme, err := repositories.GetShopThemeByID(shopThemeID)
	if err != nil {
		return errors.New("shop theme not found")
	}
	if shopTheme.ShopID != shopID {
		return errors.New("shop theme does not belong to this shop")
	}

	return repositories.PublishShopTheme(shopID, shopThemeID)
}

// DeleteShopThemeService removes a shop theme configuration
func DeleteShopThemeService(shopThemeIDStr string, sellerID primitive.ObjectID) (*mongo.DeleteResult, error) {
	shopThemeID, err := primitive.ObjectIDFromHex(shopThemeIDStr)
	if err != nil {
		return nil, errors.New("invalid shop theme ID")
	}

	// Get existing shop theme
	shopTheme, err := repositories.GetShopThemeByID(shopThemeID)
	if err != nil {
		return nil, errors.New("shop theme not found")
	}

	// Validate shop ownership
	if err := validateShopOwnership(shopTheme.ShopID, sellerID); err != nil {
		return nil, err
	}

	// Don't allow deletion of published theme
	if shopTheme.Published {
		return nil, errors.New("cannot delete published theme")
	}

	return repositories.DeleteShopTheme(shopThemeID)
}

// =============== Theme Preset Services ===============

// GetThemePresetsService retrieves presets for a theme
func GetThemePresetsService(themeIDStr string) ([]models.ThemePreset, error) {
	themeID, err := primitive.ObjectIDFromHex(themeIDStr)
	if err != nil {
		return nil, errors.New("invalid theme ID")
	}

	return repositories.GetThemePresetsByThemeID(themeID)
}

// ApplyThemePresetService applies a preset to a shop theme
func ApplyThemePresetService(shopThemeIDStr, presetIDStr string, sellerID primitive.ObjectID) (*mongo.UpdateResult, error) {
	shopThemeID, err := primitive.ObjectIDFromHex(shopThemeIDStr)
	if err != nil {
		return nil, errors.New("invalid shop theme ID")
	}

	presetID, err := primitive.ObjectIDFromHex(presetIDStr)
	if err != nil {
		return nil, errors.New("invalid preset ID")
	}

	// Get shop theme
	shopTheme, err := repositories.GetShopThemeByID(shopThemeID)
	if err != nil {
		return nil, errors.New("shop theme not found")
	}

	// Validate shop ownership
	if err := validateShopOwnership(shopTheme.ShopID, sellerID); err != nil {
		return nil, err
	}

	// Get preset
	preset, err := repositories.GetThemePresetByID(presetID)
	if err != nil {
		return nil, errors.New("preset not found")
	}

	// Verify preset belongs to the same theme
	if preset.ThemeID != shopTheme.ThemeID {
		return nil, errors.New("preset does not belong to this theme")
	}

	// Apply preset configuration
	updatedData := bson.M{
		"overrides": preset.Config,
	}

	return repositories.UpdateShopTheme(shopThemeID, updatedData)
}

// =============== Validation Functions ===============

// validateTheme validates theme data
func validateTheme(theme *models.Theme) error {
	if err := validateThemeName(theme.Name); err != nil {
		return err
	}

	if strings.TrimSpace(theme.Author) == "" {
		return errors.New("theme author is required")
	}

	if err := validateThemeConfig(&theme.Config); err != nil {
		return err
	}

	return nil
}

// validateThemeName validates theme name
func validateThemeName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("theme name is required")
	}
	if len(name) < 3 {
		return errors.New("theme name must be at least 3 characters")
	}
	if len(name) > 50 {
		return errors.New("theme name must be less than 50 characters")
	}
	return nil
}

// validateThemeConfig validates theme configuration
func validateThemeConfig(config *models.ThemeConfig) error {
	// Validate colors
	if config.Colors != nil {
		for key, value := range config.Colors {
			if key == "" || value == "" {
				return errors.New("color configuration cannot have empty keys or values")
			}
			// Basic hex color validation
			if !isValidHexColor(value) {
				return errors.New("invalid color format: " + value)
			}
		}
	}

	// Validate fonts
	if config.Fonts != nil {
		for key, value := range config.Fonts {
			if key == "" || value == "" {
				return errors.New("font configuration cannot have empty keys or values")
			}
		}
	}

	// Validate layouts
	if config.Layouts != nil {
		for key, value := range config.Layouts {
			if key == "" || value == "" {
				return errors.New("layout configuration cannot have empty keys or values")
			}
		}
	}

	// Validate sections
	for _, section := range config.Sections {
		if section.ID == "" || section.Type == "" {
			return errors.New("section must have ID and type")
		}
	}

	return nil
}

// isValidHexColor validates hex color format
func isValidHexColor(color string) bool {
	if len(color) != 7 || color[0] != '#' {
		return false
	}
	for i := 1; i < len(color); i++ {
		c := color[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// validateShopOwnership validates that a seller owns a shop
func validateShopOwnership(shopID, sellerID primitive.ObjectID) error {
	// This would typically check the shop's OwnerID against the sellerID
	// For now, we'll assume it's valid, but you should implement proper validation
	// by checking the shop document in the database
	
	// Example implementation:
	// shop, err := shopRepository.GetShopByID(shopID)
	// if err != nil {
	//     return errors.New("shop not found")
	// }
	// if shop.OwnerID != sellerID {
	//     return errors.New("unauthorized: you don't own this shop")
	// }
	
	return nil
}
