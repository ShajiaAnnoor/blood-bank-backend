package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
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
func (sc *StaticContent) FromModel(modelStaticContent *model.StaticContent) error {
	sc.Comment = modelStaticContent.StaticContent
	sc.CreatedAt = modelStaticContent.CreatedAt
	sc.UpdatedAt = modelStaticContent.UpdatedAt

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
func (c *StaticContent) ModelStaticContent() *model.StaticContent {
	sc := model.StaticContent{}
	sc.ID = c.ID.Hex()
	sc.StaticContent = c.StaticContent
	sc.UserID = c.UserID.Hex()
	sc.StaticContentID = c.StatusID.Hex()
	sc.CreatedAt = c.CreatedAt
	sc.UpdatedAt = c.UpdatedAt

	return &sc
}
