package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BloodReq holds db data type for comments
type BloodReq struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Request   string             `bson:"request"`
	UserID    primitive.ObjectID `bson:"user_id"`
	RequestID primitive.ObjectID `bson:"request_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (c *BloodReq) FromModel(modelRequest *model.BloodReq) error {
	c.Request = modelRequest.Request
	c.CreatedAt = modelRequest.CreatedAt
	c.UpdatedAt = modelRequest.UpdatedAt

	var err error
	c.RequestID, err = primitive.ObjectIDFromHex(modelRequest.RequestID)

	if err != nil {
		return err
	}

	c.UserID, err = primitive.ObjectIDFromHex(modelRequest.UserID)
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

	c.ID = id
	return nil
}

//ModelRequest converts bson to model
func (c *BloodReq) ModelRequest() *model.Request {
	bloodreq := model.Request{}
	bloodreq.ID = c.ID.Hex()
	bloodreq.Request = c.Request
	bloodreq.UserID = c.UserID.Hex()
	bloodreq.RequestID = c.RequestID.Hex()
	bloodreq.CreatedAt = c.CreatedAt
	bloodreq.UpdatedAt = c.UpdatedAt

	return &bloodreq
}
