package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Patient holds db data type for patients
type Patient struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	BloodGroup string             `bson:"blood_group"`
	District   string             `bson:"district"`
	Phone      string             `bson:"phone_number"`
	Address    string             `bson:"address"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (p *Patient) FromModel(modelPatient *model.Patient) error {
	p.CreatedAt = modelPatient.CreatedAt
	p.UpdatedAt = modelPatient.UpdatedAt
	p.Name = modelPatient.Name
	p.Address = modelPatient.Address
	p.District = modelPatient.District
	p.Phone = modelPatient.Phone
	p.Address = modelPatient.Address
	p.BloodGroup = modelPatient.BloodGroup

	var err error

	if modelPatient.ID != "" {
		p.ID, err = primitive.ObjectIDFromHex(modelPatient.ID)
	}

	if err != nil {
		return err
	}

	if modelPatient.UserID != "" {
		p.UserID, err = primitive.ObjectIDFromHex(modelPatient.UserID)
	}

	if err != nil {
		return err
	}

	return nil
}

//ModelPatient converts bson to model
func (p *Patient) ModelPatient() *model.Patient {
	patient := model.Patient{}
	patient.ID = p.ID.Hex()
	patient.UserID = p.UserID.Hex()
	patient.CreatedAt = p.CreatedAt
	patient.UpdatedAt = p.UpdatedAt
	patient.Address = p.Address
	patient.Name = p.Name
	patient.BloodGroup = p.BloodGroup
	patient.District = p.District
	patient.Phone = p.Phone
	patient.Address = p.Address

	return &patient
}
