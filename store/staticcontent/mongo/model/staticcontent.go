package model

import (
	"time"

	"gitlab.com/Aubichol/hrishi-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//StaticContent holds db data type for static contents
type StaticContent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Comment   string             `bson:"comment"`
	UserID    primitive.ObjectID `bson:"user_id"`
	StatusID  primitive.ObjectID `bson:"status_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to db data for comments
func (sc *StaticContent) FromModel(modelComment *model.StaticmodelStaticContent) error {
	sc.Comment = modelComment.Comment
	sc.CreatedAt = modelComment.CreatedAt
	sc.UpdatedAt = modelComment.UpdatedAt

	var err error
	sc.StatusID, err = primitive.ObjectIDFromHex(modelStaticContent.StatusID)

	if err != nil {
		return err
	}

	sc.UserID, err = primitive.ObjectIDFromHex(modelStaticContent.UserID)
	if err != nil {
		return err
	}

	if modelStaticContent.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelStaticContent.ID)
	if err != nil {
		return err
	}

	sc.ID = id
	return nil
}

//ModelStaticContent converts bson to model
func (c *StaticContent) ModelStaticContent() *model.Comment {
	sc := model.StaticContent{}
	sc.ID = c.ID.Hex()
	sc.ModelStaticContent = c.StaticCoModelStaticContent
	sc.UserID = c.UserID.Hex()
	sc.StaticContenID = c.StatusID.Hex()
	sc.CreatedAt = c.CreatedAt
	sc.UpdatedAt = c.UpdatedAt

	return &comment
}
