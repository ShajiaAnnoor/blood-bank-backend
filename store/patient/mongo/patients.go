package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/model"
	mongoModel "gitlab.com/Aubichol/blood-bank-backend/store/comment/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//users handles user related database queries
type patients struct {
	c *mongo.Collection
}

func (c *patients) convertData(modelNotice *model.Notice) (
	mongoNotice mongoModel.Notice,
	err error,
) {
	err = mongoNotice.FromModel(modelNotice)
	return
}

// Save saves comments from model to database
func (c *patients) Save(modelNotice *model.Notice) (string, error) {
	mongoNotice := mongoModel.Notice{}
	var err error
	mongoNotice, err = c.convertData(modelNotice)
	if err != nil {
		return "", fmt.Errorf("Could not convert model comment to mongo comment: %w", err)
	}

	if modelNotice.ID == "" {
		mongoNotice.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoNotice.ID}
	update := bson.M{"$set": mongoNotice}
	upsert := true

	_, err = c.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return mongoNotice.ID.Hex(), err
}

//FindByID finds a comment by id
func (c *patients) FindByID(id string) (*model.Patient, error) {
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

	notice := mongoModel.Patient{}
	if err := result.Decode(&notice); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return notice.ModelNotice(), nil
}

//FindByStatusID finds a comment by status id
func (c *patients) FindByNoticeID(id string, skip int64, limit int64) ([]*model.Patient, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := c.c.Find(
		context.Background(),
		filter,
		findOptions,
	)

	if err != nil {
		return nil, err
	}

	return c.cursorToPatients(cursor)
}

//CountByStatusID returns comments from status id
func (c *patients) CountByPatientID(id string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return -1, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}
	cnt, err := c.c.CountDocuments(
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
func (c *patients) FindByIDs(ids ...string) ([]*model.Patient, error) {
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

	cursor, err := c.c.Find(
		context.Background(),
		filter,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return c.cursorToPatients(cursor)
}

//Search search for users given the text, skip and limit
func (c *patients) Search(text string, skip, limit int64) ([]*model.Patient, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := c.c.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	return c.cursorToPatients(cursor)
}

//cursorToComments decodes users one by one from the search result
func (c *patients) cursorToPatients(cursor *mongo.Cursor) ([]*model.Patient, error) {
	defer cursor.Close(context.Background())
	modelPatients := []*model.Patient{}

	for cursor.Next(context.Background()) {
		patient := mongoModel.Patient{}
		if err := cursor.Decode(&patient); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelPatients = append(modelPatients, patient.ModelPatient())
	}

	return modelPatients, nil
}

//PatientssParams provides parameters for comment specific Collection
type PatientsParams struct {
	dig.In
	Collection *mongo.Collection `name:"patients"`
}

//Store provides store for comments
func Store(params PatientsParams) storepatient.Patients {
	return &patients{params.Collection}
}