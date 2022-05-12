package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Notice holds db data type for notices
type Notice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BloodGroup  string             `bson:"blood_group,omitempty"`
	Description string             `bson:"description"`
	Title       string             `bson:"title"`
	UserID      primitive.ObjectID `bson:"user_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for notices
func (n *Notice) FromModel(modelNotice *model.Notice) error {
	n.CreatedAt = modelNotice.CreatedAt
	n.UpdatedAt = modelNotice.UpdatedAt
	n.BloodGroup = modelNotice.BloodGroup
	n.Description = modelNotice.Description
	n.Title = modelNotice.Title

	var err error

	if modelNotice.ID != "" {
		n.ID, err = primitive.ObjectIDFromHex(modelNotice.ID)
	}

	if err != nil {
		return err
	}

	if modelNotice.UserID != "" {
		n.UserID, err = primitive.ObjectIDFromHex(modelNotice.UserID)
	}

	if err != nil {
		return err
	}

	if modelNotice.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelNotice.ID)
	if err != nil {
		return err
	}

	n.ID = id
	return nil
}

//ModelNotice converts bson to model
func (n *Notice) ModelNotice() *model.Notice {
	notice := model.Notice{}
	notice.ID = n.ID.Hex()
	notice.UserID = n.UserID.Hex()
	notice.CreatedAt = n.CreatedAt
	notice.UpdatedAt = n.UpdatedAt

	return &notice
}
