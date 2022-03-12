package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	mongoModel "gitlab.com/Aubichol/blood-bank-backend/store/bloodrequest/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//users handles user related database queries
type bloodrequests struct {
	c *mongo.Collection
}

func (b *bloodrequests) convertData(modelBloodRequests *model.BloodRequest) (
	mongoBloodRequests mongoModel.BloodRequest,
	err error,
) {
	err = mongoBloodRequests.FromModel(modelBloodRequests)
	return
}

// Save saves bloodrequests from model to database
func (b *bloodrequests) Save(modelBloodRequests *model.BloodRequest) (string, error) {
	mongoBloodRequests := mongoModel.BloodRequest{}
	var err error
	mongoBloodRequest, err := b.convertData(modelBloodRequests)
	if err != nil {
		return "", fmt.Errorf("Could not convert model bloodrequests to mongo bloodrequest: %w", err)
	}

	if modelBloodRequests.ID == "" {
		mongoBloodRequests.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoBloodRequest.ID}
	update := bson.M{"$set": mongoBloodRequest}
	upsert := true

	_, err = b.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoBloodRequest.ID.Hex(), err
}

//FindByID finds a blood request by id
func (b *bloodrequests) FindByID(id string) (*model.BloodRequest, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := b.c.FindOne(context.Background(), filter, &options.FindOneOptions{})
	if err := result.Err(); err != nil {
		return nil, err
	}

	bloodrequests := mongoModel.BloodRequest{}
	if err := result.Decode(&bloodrequests); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return b.ModelBloodRequests(), nil
}

//FindByBloodRequestsID finds a blood requests id
func (b *bloodrequests) FindByBloodRequestsID(id string, skip int64, limit int64) ([]*model.BloodRequest, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"blood_requests_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := b.c.Find(context.Background(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	return b.cursorToBloodRequests(cursor)
}

//CountByBloodRequestsID returns blood requests id
func (b *bloodrequests) CountByBloodRequestsID(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return -1, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}
	cnt, err := b.c.CountDocuments(context.Background(), filter, &options.CountOptions{})

	if err != nil {
		return -1, err
	}

	return cnt, nil
}

//FindByIDs returns all the blood requests from multiple blood requests
func (b *bloodrequests) FindByIDs(ids ...string) ([]*model.BloodRequest, error) {
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

	cursor, err := b.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return b.cursorToBloodRequests(cursor)
}

//Search search for blood requests given the text, skip and limit
func (b *bloodrequests) Search(text string, skip, limit int64) ([]*model.Comment, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := b.c.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	return b.cursorToBloodRequests(cursor)
}

//cursorToBloodRequests decodes blood requests one by one from the search result
func (c *bloodrequests) cursorToBloodRequests(cursor *mongo.Cursor) ([]*model.BloodRequests, error) {
	defer cursor.Close(context.Background())
	modelBloodRequests := []*model.BloodRequest{}

	for cursor.Next(context.Background()) {
		bloodreq := mongoModel.BloodRequest{}
		if err := cursor.Decode(&bloodreq); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelBloodRequests = append(modelBloodRequests, bloodreq.ModelBloodRequest())
	}

	return modelBloodRequests, nil
}

//BloodRequestsParams provides parameters for blood request specific Collection
type BloodRequestsParams struct {
	dig.In
	Collection *mongo.Collection `name:"bloodrequests"`
}

//Store provides store for blood requests
func Store(params BloodRequestsParams) storebloodrequests.BloodRequests {
	return &bloodrequests{params.Collection}
}
