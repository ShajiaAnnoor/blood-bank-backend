package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	storetoken "gitlab.com/Aubichol/blood-bank-backend/store/token"
	mongoModel "gitlab.com/Aubichol/blood-bank-backend/store/token/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//tokens handles token related database queries
type tokens struct {
	c *mongo.Collection
}

// Save saves user from model to database
func (t *tokens) Save(modelToken *model.Token) error {
	mongoToken := mongoModel.Token{}
	if err := mongoToken.FromModel(modelToken); err != nil {
		return fmt.Errorf("Could not convert model token to mongo token: %w", err)
	}

	if modelToken.ID == "" {
		mongoToken.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoToken.ID}
	update := bson.M{"$set": mongoToken}
	upsert := true

	_, err := t.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return err
}

//FindByID finds a status by id
func (t *tokens) FindByID(id string) (*model.Token, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := t.c.FindOne(context.Background(), filter, &options.FindOneOptions{})
	if err := result.Err(); err != nil {
		return nil, err
	}

	token := mongoModel.Token{}
	if err := result.Decode(&token); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return token.ModelToken(), nil
}

//FindByUser finds a status by id
func (t *tokens) FindByUser(id string) ([]*model.Token, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"user_id": objectID}
	result := t.c.FindOne(context.Background(), filter, &options.FindOneOptions{})
	if err := result.Err(); err != nil {
		return nil, err
	}

	cursor, err := t.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return t.cursorToTokens(cursor)
}

//FindByIDs returns all the users from multiple user ids
func (t *tokens) FindByIDs(ids ...string) ([]*model.Token, error) {
	objectIDs := []primitive.ObjectID{}
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("Invalid id %s : %w", id, err)
		}

		objectIDs = append(objectIDs, objectID)
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": objectIDs,
		},
	}

	cursor, err := t.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return t.cursorToTokens(cursor)
}

//Search search for users given the text, skip and limit
func (t *tokens) Search(text string, skip, limit int64) ([]*model.Token, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := t.c.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	return t.cursorToTokens(cursor)
}

//cursorToTokens decodes users one by one from the search result
func (t *tokens) cursorToTokens(cursor *mongo.Cursor) ([]*model.Token, error) {
	defer cursor.Close(context.Background())
	modelTokens := []*model.Token{}

	for cursor.Next(context.Background()) {
		token := mongoModel.Token{}
		if err := cursor.Decode(&token); err != nil {
			return nil, fmt.Errorf("could not decode data from mongo %w", err)
		}

		modelTokens = append(modelTokens, token.ModelToken())
	}

	return modelTokens, nil
}

// Params provides parameters for user specific Collection
type Params struct {
	dig.In
	Collection *mongo.Collection `name:"tokens"`
}

//Store provides store for registration tokens
func Store(params Params) storetoken.Token {
	return &tokens{params.Collection}
}
