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
type organizations struct {
	c *mongo.Collection
}

func (c *organizations) convertData(modelDonor *model.Donor) (
	mongoDonor mongoModel.Donor,
	err error,
) {
	err = mongoDonor.FromModel(modelDonor)
	return
}

// Save saves comments from model to database
func (c *organizations) Save(modelDonor *model.Donor) (string, error) {
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

//FindByID finds a organization by id
func (d *organizations) FindByID(id string) (*model.Organization, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := d.c.FindOne(
		context.Background(),
		filter,
		&options.FindOneOptions{},
	)
	if err := result.Err(); err != nil {
		return nil, err
	}

	organization := mongoModel.Organization{}
	if err := result.Decode(&organization); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return organization.ModelDonor(), nil
}

//FindByDonorID finds a donor by donor id
func (o *organizations) FindByOrganizationID(id string, skip int64, limit int64) ([]*model.Donor, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"status_id": objectID}

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"updated_at": -1})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := o.c.Find(context.Background(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	return o.cursorToOrganizations(cursor)
}

//CountByOrganizationID returns donors from donor id
func (c *organizations) CountByOrganizationID(id string) (int64, error) {
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

//FindByIDs returns all the donors from multiple donor ids
func (d *organizations) FindByIDs(ids ...string) ([]*model.Donor, error) {
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

	cursor, err := d.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return d.cursorToOrganizations(cursor)
}

//Search search for users given the text, skip and limit
func (d *organizations) Search(text string, skip, limit int64) ([]*model.Organization, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := d.c.Find(
		context.Background(),
		filter,
		&options.FindOptions{
			Skip:  &skip,
			Limit: &limit,
		}
	)
	if err != nil {
		return nil, err
	}

	return c.cursorToOrganizations(cursor)
}

//cursorToDonors decodes organizations one by one from the search result
func (d *donors) cursorToOrganizations(cursor *mongo.Cursor) ([]*model.Organization, error) {
	defer cursor.Close(context.Background())
	modelDonors := []*model.Donor{}

	for cursor.Next(context.Background()) {
		donor := mongoModel.Organization{}
		if err := cursor.Decode(&donor); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongo %w", err)
		}

		modelDonors = append(modelOrganizations, donor.ModelOrganization())
	}

	return modelOrganizations, nil
}

//OrganizationsParams provides parameters for organization specific Collection
type OrganizationsParams struct {
	dig.In
	Collection *mongo.Collection `name:"organizations"`
}

//Store provides store for organizations
func Store(params OrganizationsParams) storedonor.Organizations {
	return &Organizations{params.Collection}
}
