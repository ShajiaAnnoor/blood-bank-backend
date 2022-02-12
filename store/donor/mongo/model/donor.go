package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Donor holds db data type for donors
type Donor struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Comment   string             `bson:"comment"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (c *Donor) FromModel(modelDonor *model.Donor) error {
	c.Comment = modelComment.Comment
	c.CreatedAt = modelComment.CreatedAt
	c.UpdatedAt = modelComment.UpdatedAt

	var err error
	c.StatusID, err = primitive.ObjectIDFromHex(modelComment.StatusID)

	if err != nil {
		return err
	}

	c.UserID, err = primitive.ObjectIDFromHex(modelComment.UserID)
	if err != nil {
		return err
	}

	if modelComment.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelComment.ID)
	if err != nil {
		return err
	}

	c.ID = id
	return nil
}

//ModelComment converts bson to model
func (c *Donor) ModelDonor() *model.Donor {
	donor := model.Donor{}
	donor.ID = c.ID.Hex()
	donor.Comment = c.Comment
	donor.UserID = c.UserID.Hex()
	donor.StatusID = c.StatusID.Hex()
	donor.CreatedAt = c.CreatedAt
	donor.UpdatedAt = c.UpdatedAt

	return &donor
}
