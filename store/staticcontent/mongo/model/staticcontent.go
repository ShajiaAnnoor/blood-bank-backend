package model

import (
	"time"

	"gitlab.com/Aubichol/hrishi-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Comment holds db data type for comments
type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Comment   string             `bson:"comment"`
	UserID    primitive.ObjectID `bson:"user_id"`
	StatusID  primitive.ObjectID `bson:"status_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (c *Comment) FromModel(modelComment *model.Comment) error {
	c.Comment = modelComment.Comment
	c.CreatedAt = modelComment.CreatedAt
	c.UpdatedAt = modelComment.UpdatedAt

	var err error
	c.StatusID, err = primitive.ObjectIDFromHex(modelComment.StatusID)

	if err != nil {
		return err
	}

	c.UserID, err = primitive.ObjectIDFromHex(modelComment.UserID)
	if err != nil {
		return err
	}

	if modelComment.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelComment.ID)
	if err != nil {
		return err
	}

	c.ID = id
	return nil
}

//ModelComment converts bson to model
func (c *Comment) ModelComment() *model.Comment {
	comment := model.Comment{}
	comment.ID = c.ID.Hex()
	comment.Comment = c.Comment
	comment.UserID = c.UserID.Hex()
	comment.StatusID = c.StatusID.Hex()
	comment.CreatedAt = c.CreatedAt
	comment.UpdatedAt = c.UpdatedAt

	return &comment
}
