package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BloodReq holds db data type for blood requests
type BloodRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Request   string             `bson:"request"`
	UserID    primitive.ObjectID `bson:"user_id"`
	RequestID primitive.ObjectID `bson:"request_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for blood requests
func (b *BloodRequest) FromModel(modelRequest *model.BloodRequest) error {
	b.Request = modelRequest.Request
	b.CreatedAt = modelRequest.CreatedAt
	b.UpdatedAt = modelRequest.UpdatedAt

	var err error
	b.RequestID, err = primitive.ObjectIDFromHex(modelRequest.ID)

	if err != nil {
		return err
	}

	b.UserID, err = primitive.ObjectIDFromHex(modelRequest.UserID)
	if err != nil {
		return err
	}

	if modelRequest.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelRequest.ID)
	if err != nil {
		return err
	}

	b.ID = id
	return nil
}

//ModelRequest converts bson to model
func (b *BloodRequest) ModelBloodRequest() *model.BloodRequest {
	bloodreq := model.BloodRequest{}
	bloodreq.ID = b.ID.Hex()
	bloodreq.Request = b.Request
	bloodreq.UserID = b.UserID.Hex()
	bloodreq.ID = b.RequestID.Hex()
	bloodreq.CreatedAt = b.CreatedAt
	bloodreq.UpdatedAt = b.UpdatedAt

	return &bloodreq
}
