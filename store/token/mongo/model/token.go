package model

import (
	"time"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Token defines mongodb data type for Token
type Token struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Token     string             `bson:"token"`
	UserID    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

//FromModel converts model data to bson data
func (t *Token) FromModel(modelToken *model.Token) error {
	t.Token = modelToken.Token
	t.CreatedAt = modelToken.CreatedAt
	t.UpdatedAt = modelToken.UpdatedAt

	var err error

	t.UserID, err = primitive.ObjectIDFromHex(modelToken.UserID)

	if err != nil {
		return err
	}

	return nil
}

//ModelToken converts bson to model for token
func (t *Token) ModelToken() *model.Token {
	token := model.Token{}
	token.ID = t.ID.Hex()
	token.Token = t.Token
	token.UserID = t.ID.Hex()
	token.CreatedAt = t.CreatedAt
	token.UpdatedAt = t.UpdatedAt

	return &token
}
