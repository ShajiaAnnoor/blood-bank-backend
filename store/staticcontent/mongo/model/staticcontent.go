package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//StaticContent holds db data type for static contents
type StaticContent struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Text      string             `bson:"text,omitempty"`
	IsDeleted bool               `bson:"is_deleted"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

//FromModel converts model data to db data for comments
func (sc *StaticContent) FromModel(modelStaticContent *model.StaticContent) error {
	sc.CreatedAt = modelStaticContent.CreatedAt
	sc.UpdatedAt = modelStaticContent.UpdatedAt
	sc.Text = modelStaticContent.Text
	sc.IsDeleted = modelStaticContent.IsDeleted

	var err error
	if modelStaticContent.ID != "" {
		sc.ID, err = primitive.ObjectIDFromHex(modelStaticContent.ID)
	}

	if err != nil {
		return err
	}

	if modelStaticContent.UserID != "" {
		sc.UserID, err = primitive.ObjectIDFromHex(modelStaticContent.UserID)
	}

	if err != nil {
		return err
	}

	return nil
}

//ModelStaticContent converts bson to model
func (c *StaticContent) ModelStaticContent() *model.StaticContent {
	sc := model.StaticContent{}
	sc.ID = c.ID.Hex()
	sc.Text = c.Text
	sc.UserID = c.UserID.Hex()
	sc.CreatedAt = c.CreatedAt
	sc.UpdatedAt = c.UpdatedAt

	return &sc
}
