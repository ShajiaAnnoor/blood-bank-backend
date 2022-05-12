package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Comment holds db data type for comments
type Organization struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description,omitempty"`
	Phone       string             `bson:"phone_number,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Address     string             `bson:"address,omitempty"`
	District    string             `bson:"district,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
}

//FromModel converts model data to db data for comments
func (o *Organization) FromModel(modelOrganization *model.Organization) error {
	o.CreatedAt = modelOrganization.CreatedAt
	o.UpdatedAt = modelOrganization.UpdatedAt
	o.Name = modelOrganization.Name
	o.Address = modelOrganization.Address
	o.Description = modelOrganization.Description
	o.Phone = modelOrganization.Phone
	o.District = modelOrganization.District

	var err error

	if modelOrganization.ID != "" {
		o.ID, err = primitive.ObjectIDFromHex(modelOrganization.ID)
	}

	if err != nil {
		return err
	}

	if modelOrganization.UserID != "" {
		o.UserID, err = primitive.ObjectIDFromHex(modelOrganization.UserID)
	}

	if err != nil {
		return err
	}

	return nil
}

//ModelOrganization converts bson to model
func (o *Organization) ModelOrganization() *model.Organization {
	organization := model.Organization{}
	organization.ID = o.ID.Hex()
	organization.UserID = o.UserID.Hex()
	organization.CreatedAt = o.CreatedAt
	organization.UpdatedAt = o.UpdatedAt
	organization.Name = o.Name
	organization.Address = o.Address
	organization.Description = o.Description
	organization.Phone = o.Phone

	return &organization
}
