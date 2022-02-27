package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	mongoModel "gitlab.com/Aubichol/blood-bank-backend/store/notice/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//notices handles notice related database queries
type notices struct {
	c *mongo.Collection
}

func (n *notices) convertData(modelNotice *model.Notice) (
	mongoNotice mongoModel.Notice,
	err error,
) {
	err = mongoNotice.FromModel(modelNotice)
	return
}

// Save saves notices from model to database
func (n *notices) Save(modelNotice *model.Notice) (string, error) {
	mongoNotice := mongoModel.Notice{}
	var err error
	mongoNotice, err = c.convertData(modelNotice)
	if err != nil {
		return "", fmt.Errorf("Could not convert model notice to mongo notice: %w", err)
	}

	if modelNotice.ID == "" {
		mongoNotice.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoNotice.ID}
	update := bson.M{"$set": mongoNotice}
	upsert := true

	_, err = n.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoNotice.ID.Hex(), err
}

//FindByID finds a notice by id
func (n *notices) FindByID(id string) (*model.Notice, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := n.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)
	if err := result.Err(); err != nil {
		return nil, err
	}

	notice := mongoModel.Notice{}
	if err := result.Decode(&notice); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return notice.ModelNotice(), nil
}

//FindByStatusID finds a comment by status id
func (n *notices) FindByNoticeID(id string, skip int64, limit int64) ([]*model.Notice, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"notice_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := n.c.Find(
		context.Background(),
		filter,
		findOptions,
	)

	if err != nil {
		return nil, err
	}

	return n.cursorToNotices(cursor)
}

//CountByNoticeID returns notices from notice id
func (n *notices) CountByNoticeID(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return -1, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"notice_id": objectID}
	cnt, err := n.c.CountDocuments(
		context.Background(), 
		filter, 
		&options.CountOptions{}
	)

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

//FindByIDs returns all the users from multiple user ids
func (n *notices) FindByIDs(ids ...string) ([]*model.Notice, error) {
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

	cursor, err := n.c.Find(
		context.Background(),
		filter,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return n.cursorToNotices(cursor)
}

//Search search for notices given the text, skip and limit
func (n *notices) Search(text string, skip, limit int64) ([]*model.Notice, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := n.c.Find(
		context.Background(),
		filter,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		})
	if err != nil {
		return nil, err
	}

	return n.cursorToNotices(cursor)
}

//cursorToNotices decodes notices one by one from the search result
func (n *notices) cursorToNotices(cursor *mongo.Cursor) ([]*model.Notice, error) {
	defer cursor.Close(context.Background())
	modelNotices := []*model.Notice{}

	for cursor.Next(context.Background()) {
		notice := mongoModel.Notice{}
		if err := cursor.Decode(&notice); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelNotices = append(modelNotices, notice.ModelNotice())
	}

	return modelNotices, nil
}

//NoticesParams provides parameters for notice specific Collection
type NoticesParams struct {
	dig.In
	Collection *mongo.Collection `name:"notices"`
}

//Store provides store for notices
func Store(params NoticesParams) storenotice.Notices {
	return &notices{params.Collection}
}
