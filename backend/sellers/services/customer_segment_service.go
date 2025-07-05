package services

import (
	"errors"

	"github.com/Endale2/DRPS/sellers/models"
	"github.com/Endale2/DRPS/sellers/repositories"
	sharedShopService "github.com/Endale2/DRPS/shared/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerSegmentService encapsulates business logic for customer segments
type CustomerSegmentService struct{}

// NewCustomerSegmentService returns a new service instance
func NewCustomerSegmentService() *CustomerSegmentService {
	return &CustomerSegmentService{}
}

// CreateCustomerSegmentService creates a new customer segment for a given shop
func (svc *CustomerSegmentService) CreateCustomerSegmentService(
	shopID primitive.ObjectID,
	sellerID primitive.ObjectID,
	name, description, color string,
) (*models.CustomerSegment, error) {
	// Verify this shop exists and belongs to seller
	shop, err := sharedShopService.GetShopByIDService(shopID.Hex())
	if err != nil {
		return nil, err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return nil, errors.New("not authorized to create segments for this shop")
	}

	// Build the CustomerSegment model
	segment := &models.CustomerSegment{
		ShopID:      shopID,
		Name:        name,
		Description: description,
		Color:       color,
		CustomerIDs: []primitive.ObjectID{},
	}

	_, err = repositories.CreateCustomerSegment(segment)
	if err != nil {
		return nil, err
	}
	return segment, nil
}

// GetCustomerSegmentsByShopService returns all customer segments for a given shop
func (svc *CustomerSegmentService) GetCustomerSegmentsByShopService(shopID primitive.ObjectID) ([]models.CustomerSegment, error) {
	return repositories.GetCustomerSegmentsByShop(shopID)
}

// UpdateCustomerSegmentService updates a customer segment's fields
func (svc *CustomerSegmentService) UpdateCustomerSegmentService(
	segmentID primitive.ObjectID,
	sellerID primitive.ObjectID,
	updates bson.M,
) error {
	// Verify segment exists and belongs to seller's shop
	segment, err := repositories.GetCustomerSegmentByID(segmentID)
	if err != nil {
		return err
	}
	if segment == nil {
		return errors.New("segment not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(segment.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to update this segment")
	}

	_, err = repositories.UpdateCustomerSegment(segmentID, updates)
	return err
}

// DeleteCustomerSegmentService removes a customer segment
func (svc *CustomerSegmentService) DeleteCustomerSegmentService(
	segmentID primitive.ObjectID,
	sellerID primitive.ObjectID,
) error {
	// Verify segment exists and belongs to seller's shop
	segment, err := repositories.GetCustomerSegmentByID(segmentID)
	if err != nil {
		return err
	}
	if segment == nil {
		return errors.New("segment not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(segment.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to delete this segment")
	}

	_, err = repositories.DeleteCustomerSegment(segmentID)
	return err
}

// AddCustomerToSegmentService adds a customer to a segment
func (svc *CustomerSegmentService) AddCustomerToSegmentService(
	segmentID primitive.ObjectID,
	customerID primitive.ObjectID,
	sellerID primitive.ObjectID,
) error {
	// Verify segment exists and belongs to seller's shop
	segment, err := repositories.GetCustomerSegmentByID(segmentID)
	if err != nil {
		return err
	}
	if segment == nil {
		return errors.New("segment not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(segment.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this segment")
	}

	// Verify customer is linked to this shop
	linked, err := sharedShopService.IsCustomerLinked(shop.ID, customerID)
	if err != nil {
		return err
	}
	if !linked {
		return errors.New("customer is not linked to this shop")
	}

	_, err = repositories.AddCustomerToSegment(segmentID, customerID)
	return err
}

// RemoveCustomerFromSegmentService removes a customer from a segment
func (svc *CustomerSegmentService) RemoveCustomerFromSegmentService(
	segmentID primitive.ObjectID,
	customerID primitive.ObjectID,
	sellerID primitive.ObjectID,
) error {
	// Verify segment exists and belongs to seller's shop
	segment, err := repositories.GetCustomerSegmentByID(segmentID)
	if err != nil {
		return err
	}
	if segment == nil {
		return errors.New("segment not found")
	}

	// Verify shop ownership
	shop, err := sharedShopService.GetShopByIDService(segment.ShopID.Hex())
	if err != nil {
		return err
	}
	if shop == nil || shop.OwnerID != sellerID {
		return errors.New("not authorized to modify this segment")
	}

	_, err = repositories.RemoveCustomerFromSegment(segmentID, customerID)
	return err
}

// GetSegmentsByCustomerService returns all segments that contain a specific customer
func (svc *CustomerSegmentService) GetSegmentsByCustomerService(
	shopID primitive.ObjectID,
	customerID primitive.ObjectID,
) ([]models.CustomerSegment, error) {
	return repositories.GetSegmentsByCustomer(shopID, customerID)
}
