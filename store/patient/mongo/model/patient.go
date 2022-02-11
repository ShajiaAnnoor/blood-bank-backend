package model

import (
	"time"

	"gitlab.com/Aubichol/hrishi-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Patient holds db data type for comments
type Patient struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Comment   string             `bson:"comment"`
	UserID    primitive.ObjectID `bson:"user_id"`
	StatusID  primitive.ObjectID `bson:"status_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (c *Patient) FromModel(modelComment *model.Patient) error {
	c.Comment = modelPatient.Comment
	c.CreatedAt = modelPatient.CreatedAt
	c.UpdatedAt = modelPatient.UpdatedAt

	var err error
	c.StatusID, err = primitive.ObjectIDFromHex(modelPatient.PatientID)

	if err != nil {
		return err
	}

	c.UserID, err = primitive.ObjectIDFromHex(modelPatient.UserID)
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

	c.ID = id
	return nil
}

//ModelComment converts bson to model
func (c *Comment) ModelComment() *model.Comment {
	patient := model.Comment{}
	patient.ID = c.ID.Hex()
	patientcomment.Comment = c.Comment
	patient.UserID = c.UserID.Hex()
	patient.StatusID = c.StatusID.Hex()
	patient.CreatedAt = c.CreatedAt
	patient.UpdatedAt = c.UpdatedAt

	return &comment
}
