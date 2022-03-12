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

//FromModel converts model data to db data for donors
func (d *Donor) FromModel(modelDonor *model.Donor) error {
	//	d.Donor = modelDonor.Donor
	d.CreatedAt = modelDonor.CreatedAt
	d.UpdatedAt = modelDonor.UpdatedAt

	var err error
	d.ID, err = primitive.ObjectIDFromHex(modelDonor.ID)

	if err != nil {
		return err
	}

	d.UserID, err = primitive.ObjectIDFromHex(modelDonor.UserID)
	if err != nil {
		return err
	}

	if modelDonor.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelDonor.ID)
	if err != nil {
		return err
	}

	d.ID = id
	return nil
}

//ModelDonor converts bson to model
func (d *Donor) ModelDonor() *model.Donor {
	donor := model.Donor{}
	donor.ID = d.ID.Hex()
	//	donor.Comment = d.Comment
	donor.UserID = d.UserID.Hex()
	donor.ID = d.ID.Hex()
	donor.CreatedAt = d.CreatedAt
	donor.UpdatedAt = d.UpdatedAt

	return &donor
}
