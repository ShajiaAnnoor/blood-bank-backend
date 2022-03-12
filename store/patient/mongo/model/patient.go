package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Patient holds db data type for patients
type Patient struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Comment   string             `bson:"comment"`
	UserID    primitive.ObjectID `bson:"user_id"`
	StatusID  primitive.ObjectID `bson:"status_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (p *Patient) FromModel(modelPatient *model.Patient) error {
	p.Comment = modelPatient.Patient
	p.CreatedAt = modelPatient.CreatedAt
	p.UpdatedAt = modelPatient.UpdatedAt

	var err error
	p.ID, err = primitive.ObjectIDFromHex(modelPatient.ID)

	if err != nil {
		return err
	}

	p.UserID, err = primitive.ObjectIDFromHex(modelPatient.UserID)
	if err != nil {
		return err
	}

	if modelPatient.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelPatient.ID)
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

//ModelPatient converts bson to model
func (p *Patient) ModelPatient() *model.Patient {
	patient := model.Patient{}
	patient.ID = p.ID.Hex()
	patient.Patient = p.Comment
	patient.UserID = p.UserID.Hex()
	//	patient.StatusID = p.StatusID.Hex()
	patient.CreatedAt = p.CreatedAt
	patient.UpdatedAt = p.UpdatedAt

	return &patient
}
