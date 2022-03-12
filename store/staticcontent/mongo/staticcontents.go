package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	storestaticcontent "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent"
	mongoModel "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//static cobtents handles static contents related database queries
type staticcontents struct {
	c *mongo.Collection
}

func (sc *staticcontents) convertData(modelStaticContent *model.StaticContent) (
	mongoStaticContent mongoModel.StaticContent,
	err error,
) {
	err = mongoStaticContent.FromModel(modelStaticContent)
	return
}

// Save saves staticcontents from model to database
func (sc *staticcontents) Save(modelStaticContent *model.StaticContent) (string, error) {
	mongoStaticContent := mongoModel.StaticContent{}
	var err error
	mongoStaticContent, err = sc.convertData(modelStaticContent)
	if err != nil {
		return "", fmt.Errorf("Could not convert model static contents to mongo static contents: %w", err)
	}

	if modelStaticContent.ID == "" {
		mongoStaticContent.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoStaticContent.ID}
	update := bson.M{"$set": mongoStaticContent}
	upsert := true

	_, err = sc.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoStaticContent.ID.Hex(), err
}

//FindByID finds a static content by id
func (sc *staticcontents) FindByID(id string) (*model.StaticContent, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := sc.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)
	if err := result.Err(); err != nil {
		return nil, err
	}

	staticcontent := mongoModel.StaticContent{}
	if err := result.Decode(&staticcontent); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return staticcontent.ModelStaticContent(), nil
}

//FindByStatusID finds a static content by id
func (sc *staticcontents) FindByStaticContentID(id string, skip int64, limit int64) ([]*model.StaticContent, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := sc.c.Find(
		context.Background(),
		filter,
		findOptions,
	)

	if err != nil {
		return nil, err
	}

	return sc.cursorToStaticContents(cursor)
}

//CountByStaticContentID returns comments from status id
func (sc *staticcontents) CountByStaticContentID(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return -1, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}
	cnt, err := sc.c.CountDocuments(
		context.Background(),
		filter,
		&options.CountOptions{},
	)

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

//FindByIDs returns all the users from multiple user ids
func (sc *staticcontents) FindByIDs(ids ...string) ([]*model.StaticContent, error) {
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

	cursor, err := sc.c.Find(
		context.Background(),
		filter,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return sc.cursorToStaticContents(cursor)
}

//Search search for users given the text, skip and limit
func (sc *staticcontents) Search(text string, skip, limit int64) ([]*model.StaticContent, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := sc.c.Find(
		context.Background(),
		filter,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	if err != nil {
		return nil, err
	}

	return sc.cursorToStaticContents(cursor)
}

//cursorToStaticContents decodes static contents one by one from the search result
func (s *staticcontents) cursorToStaticContents(cursor *mongo.Cursor) ([]*model.StaticContent, error) {
	defer cursor.Close(context.Background())
	modelStaticContents := []*model.StaticContent{}

	for cursor.Next(context.Background()) {
		staticcontent := mongoModel.StaticContent{}
		if err := cursor.Decode(&staticcontent); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelStaticContents = append(modelStaticContents, staticcontent.ModelStaticContent())
	}

	return modelStaticContents, nil
}

//StaticContentsParams provides parameters for static content specific Collection
type StaticContentsParams struct {
	dig.In
	Collection *mongo.Collection `name:"staticcontents"`
}

//Store provides store for comments
func Store(params StaticContentsParams) storestaticcontent.StaticContents {
	return &staticcontents{params.Collection}
}
