package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/hrishi-backend/model"
	storecomment "gitlab.com/Aubichol/hrishi-backend/store/comment"
	mongoModel "gitlab.com/Aubichol/hrishi-backend/store/comment/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//users handles user related database queries
type comments struct {
	c *mongo.Collection
}

func (c *comments) convertData(modelComment *model.Comment) (
	mongoComment mongoModel.Comment,
	err error,
) {
	err = mongoComment.FromModel(modelComment)
	return
}

// Save saves comments from model to database
func (c *comments) Save(modelComment *model.Comment) (string, error) {
	mongoComment := mongoModel.Comment{}
	var err error
	mongoComment, err = c.convertData(modelComment)
	if err != nil {
		return "", fmt.Errorf("Could not convert model comment to mongo comment: %w", err)
	}

	if modelComment.ID == "" {
		mongoComment.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoComment.ID}
	update := bson.M{"$set": mongoComment}
	upsert := true

	_, err = c.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoComment.ID.Hex(), err
}

//FindByID finds a comment by id
func (c *comments) FindByID(id string) (*model.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := c.c.FindOne(context.Background(), filter, &options.FindOneOptions{})
	if err := result.Err(); err != nil {
		return nil, err
	}

	comment := mongoModel.Comment{}
	if err := result.Decode(&comment); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return comment.ModelComment(), nil
}

//FindByStatusID finds a comment by status id
func (c *comments) FindByStatusID(id string, skip int64, limit int64) ([]*model.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := c.c.Find(context.Background(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	return c.cursorToComments(cursor)
}

//CountByStatusID returns comments from status id
func (c *comments) CountByStatusID(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return -1, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}
	cnt, err := c.c.CountDocuments(context.Background(), filter, &options.CountOptions{})

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

//FindByIDs returns all the users from multiple user ids
func (c *comments) FindByIDs(ids ...string) ([]*model.Comment, error) {
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

	cursor, err := c.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return c.cursorToComments(cursor)
}

//Search search for users given the text, skip and limit
func (c *comments) Search(text string, skip, limit int64) ([]*model.Comment, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := c.c.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	return c.cursorToComments(cursor)
}

//cursorToComments decodes users one by one from the search result
func (c *comments) cursorToComments(cursor *mongo.Cursor) ([]*model.Comment, error) {
	defer cursor.Close(context.Background())
	modelComments := []*model.Comment{}

	for cursor.Next(context.Background()) {
		comment := mongoModel.Comment{}
		if err := cursor.Decode(&comment); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelComments = append(modelComments, comment.ModelComment())
	}

	return modelComments, nil
}

//CommentsParams provides parameters for comment specific Collection
type CommentsParams struct {
	dig.In
	Collection *mongo.Collection `name:"comments"`
}

//Store provides store for comments
func Store(params CommentsParams) storecomment.Comments {
	return &comments{params.Collection}
}
