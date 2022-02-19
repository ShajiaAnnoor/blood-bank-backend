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
func (c *Organization) FromModel(modelOrganization *model.Organization) error {
	c.Organization = modelOrganization.Organization
	c.CreatedAt = modelOrganization.CreatedAt
	c.UpdatedAt = modelOrganization.UpdatedAt

	var err error
	c.StatusID, err = primitive.ObjectIDFromHex(modelOrganization.StatusID)

	if err != nil {
		return err
	}

	c.UserID, err = primitive.ObjectIDFromHex(modelOrganization.UserID)
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

	c.ID = id
	return nil
}

//ModelOrganization converts bson to model
func (c *Organization) ModelOrganization() *model.Organization {
	organization := model.Organization{}
	organization.ID = c.ID.Hex()
	organization.Comment = c.Organization
	organization.UserID = c.UserID.Hex()
	organization.StatusID = c.StatusID.Hex()
	organization.CreatedAt = c.CreatedAt
	organization.UpdatedAt = c.UpdatedAt

	return &organization
}
