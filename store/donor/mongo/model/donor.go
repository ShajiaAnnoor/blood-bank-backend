package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Donor holds db data type for donors
type Donor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Phone        string             `bson:"phone_number"`
	District     string             `bson:"district"`
	Address      string             `bson:"address"`
	Availability bool               `bson:"availability"`
	TimesDonated int                `bson:"times_donated"`
	BloodGroup   string             `bson:"blood_group"`
	UserID       primitive.ObjectID `bson:"user_id"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for donors
func (d *Donor) FromModel(modelDonor *model.Donor) error {
	//	d.Donor = modelDonor.Donor
	d.CreatedAt = modelDonor.CreatedAt
	d.UpdatedAt = modelDonor.UpdatedAt
	d.Name = modelDonor.Name
	d.Phone = modelDonor.Phone
	d.District = modelDonor.District
	d.Address = modelDonor.Address
	d.Availability = modelDonor.Availability
	d.TimesDonated = modelDonor.TimesDonated
	d.BloodGroup = modelDonor.BloodGroup

	var err error

	if modelDonor.ID != "" {
		d.ID, err = primitive.ObjectIDFromHex(modelDonor.ID)
	} else {
		d.ID = primitive.NewObjectID()
	}

	if err != nil {
		return err
	}

	d.UserID, err = primitive.ObjectIDFromHex(modelDonor.UserID)

	if err != nil {
		return err
	}

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
