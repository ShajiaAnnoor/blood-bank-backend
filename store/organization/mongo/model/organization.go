package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Comment holds db data type for comments
type Organization struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Organization string             `bson:"organization"`
	UserID       primitive.ObjectID `bson:"user_id"`
	StatusID     primitive.ObjectID `bson:"status_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (o *Organization) FromModel(modelOrganization *model.Organization) error {
	o.Organization = modelOrganization.Organization
	o.CreatedAt = modelOrganization.CreatedAt
	o.UpdatedAt = modelOrganization.UpdatedAt

	var err error
	o.StatusID, err = primitive.ObjectIDFromHex(modelOrganization.StatusID)

	if err != nil {
		return err
	}

	o.UserID, err = primitive.ObjectIDFromHex(modelOrganization.UserID)
	if err != nil {
		return err
	}

	if modelOrganization.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelOrganization.ID)
	if err != nil {
		return err
	}

	o.ID = id
	return nil
}

//ModelOrganization converts bson to model
func (o *Organization) ModelOrganization() *model.Organization {
	organization := model.Organization{}
	organization.ID = o.ID.Hex()
	organization.Comment = o.Organization
	organization.UserID = o.UserID.Hex()
	organization.StatusID = o.StatusID.Hex()
	organization.CreatedAt = o.CreatedAt
	organization.UpdatedAt = o.UpdatedAt

	return &organization
}
