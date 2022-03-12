package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Notice holds db data type for notices
type Notice struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Notice    string             `bson:"notice"`
	UserID    primitive.ObjectID `bson:"user_id"`
	NoticeID  primitive.ObjectID `bson:"status_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for notices
func (n *Notice) FromModel(modelNotice *model.Notice) error {
	n.Notice = modelNotice.Notice
	n.CreatedAt = modelNotice.CreatedAt
	n.UpdatedAt = modelNotice.UpdatedAt

	var err error
	n.ID, err = primitive.ObjectIDFromHex(modelNotice.ID)

	if err != nil {
		return err
	}

	n.UserID, err = primitive.ObjectIDFromHex(modelNotice.UserID)
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
	notice.Notice = n.Notice
	notice.UserID = n.UserID.Hex()
	//	notice.StatusID = n.StatusID.Hex()
	notice.CreatedAt = n.CreatedAt
	notice.UpdatedAt = n.UpdatedAt

	return &notice
}
