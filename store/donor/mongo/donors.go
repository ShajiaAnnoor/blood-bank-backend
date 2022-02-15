package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	mongoModel "gitlab.com/Aubichol/blood-bank-backend/store/donor/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//users handles user related database queries
type donors struct {
	c *mongo.Collection
}

func (c *donors) convertData(modelDonor *model.Donor) (
	mongoDonor mongoModel.Donor,
	err error,
) {
	err = mongoDonor.FromModel(modelDonor)
	return
}

// Save saves comments from model to database
func (c *donors) Save(modelDonor *model.Donor) (string, error) {
	mongoDonor := mongoModel.Donor{}
	var err error
	mongoDonor, err = c.convertData(modelDonor)
	if err != nil {
		return "", fmt.Errorf("Could not convert model donor to mongo donor: %w", err)
	}

	if modelDonor.ID == "" {
		mongoDonor.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoDonor.ID}
	update := bson.M{"$set": mongoDonor}
	upsert := true

	_, err = c.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoDonor.ID.Hex(), err
}

//FindByID finds a donor by id
func (c *donors) FindByID(id string) (*model.Donor, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := c.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)
	if err := result.Err(); err != nil {
		return nil, err
	}

	donor := mongoModel.Donor{}
	if err := result.Decode(&donor); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return donor.ModelDonor(), nil
}

//FindByStatusID finds a comment by status id
func (c *donors) FindByDonorID(id string, skip int64, limit int64) ([]*model.Donor, error) {
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

	return c.cursorToDonors(cursor)
}

//CountByStatusID returns comments from status id
func (c *donors) CountByDonorID(id string) (int64, error) {
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
func (c *donors) FindByIDs(ids ...string) ([]*model.Comment, error) {
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

	return c.cursorToDonors(cursor)
}

//Search search for users given the text, skip and limit
func (c *donors) Search(text string, skip, limit int64) ([]*model.Comment, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := c.c.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	return c.cursorToDonors(cursor)
}

//cursorToComments decodes users one by one from the search result
func (c *donors) cursorToDonors(cursor *mongo.Cursor) ([]*model.Donor, error) {
	defer cursor.Close(context.Background())
	modelDonors := []*model.Donor{}

	for cursor.Next(context.Background()) {
		donor := mongoModel.Donor{}
		if err := cursor.Decode(&donor); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDonors = append(modelDonors, donor.ModelDonor())
	}

	return modelDonors, nil
}

//DonorsParams provides parameters for comment specific Collection
type DonorsParams struct {
	dig.In
	Collection *mongo.Collection `name:"donors"`
}

//Store provides store for comments
func Store(params DonorsParams) storedonor.Donors {
	return &donors{params.Collection}
}
