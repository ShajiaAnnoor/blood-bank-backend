package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Donor holds db data type for donors
type Donor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Phone        string             `bson:"phone_number,omitempty"`
	District     string             `bson:"district,omitempty"`
	Address      string             `bson:"address,omitempty"`
	Availability bool               `bson:"availability,omitempty"`
	TimesDonated int                `bson:"times_donated,omitempty"`
	BloodGroup   string             `bson:"blood_group,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
	IsDeleted    bool               `bson:"is_deleted,omitempty"`
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
	d.IsDeleted = modelDonor.IsDeleted

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
	donor.UserID = d.UserID.Hex()
	donor.CreatedAt = d.CreatedAt
	donor.UpdatedAt = d.UpdatedAt
	donor.Address = d.Address
	donor.Availability = d.Availability
	donor.BloodGroup = d.BloodGroup
	donor.District = d.District
	donor.Name = d.Name
	donor.Phone = d.Phone
	donor.TimesDonated = d.TimesDonated

	return &donor
}
