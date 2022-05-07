package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BloodReq holds db data type for blood requests
type BloodRequest struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Request string             `bson:"request"`
	UserID  primitive.ObjectID `bson:"user_id"`
	//	RequestID  primitive.ObjectID `bson:"request_id"`
	BloodGroup string    `bson:"blood_group"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
	IsDeleted  bool      `bson:"is_deleted"`
}

//FromModel converts model data to db data for blood requests
func (b *BloodRequest) FromModel(modelRequest *model.BloodRequest) error {
	b.Request = modelRequest.Request
	b.CreatedAt = modelRequest.CreatedAt
	b.UpdatedAt = modelRequest.UpdatedAt
	b.BloodGroup = modelRequest.BloodGroup
	b.IsDeleted = modelRequest.IsDeleted

	var err error

	b.UserID, err = primitive.ObjectIDFromHex(modelRequest.UserID)

	if err != nil {
		return err
	}

	if modelRequest.ID != "" {
		b.ID, err = primitive.ObjectIDFromHex(modelRequest.ID)
	}

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
	bloodreq.CreatedAt = b.CreatedAt
	bloodreq.UpdatedAt = b.UpdatedAt
	bloodreq.BloodGroup = b.BloodGroup

	return &bloodreq
}
