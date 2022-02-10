package model

import (
	"time"

	"gitlab.com/Aubichol/hrishi-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Notice holds db data type for comments
type Notice struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Notice    string             `bson:"notice"`
	UserID    primitive.ObjectID `bson:"user_id"`
	NoticeID  primitive.ObjectID `bson:"status_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (c *Notice) FromModel(modelNotice *model.Notice) error {
	c.Notice = modelNotice.Notice
	c.CreatedAt = modelNotice.CreatedAt
	c.UpdatedAt = modelNotice.UpdatedAt

	var err error
	c.StatusID, err = primitive.ObjectIDFromHex(modelNotice.StatusID)

	if err != nil {
		return err
	}

	c.UserID, err = primitive.ObjectIDFromHex(modelNotice.UserID)
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

	c.ID = id
	return nil
}

//ModelComment converts bson to model
func (c *Notice) ModelNotice() *model.Notice {
	notice := model.Notice{}
	notice.ID = c.ID.Hex()
	notice.Notice = c.Notice
	notice.UserID = c.UserID.Hex()
	notice.StatusID = c.StatusID.Hex()
	notice.CreatedAt = c.CreatedAt
	notice.UpdatedAt = c.UpdatedAt

	return &notice
}
