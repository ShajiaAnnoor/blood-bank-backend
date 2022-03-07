package mongo

import (
	"context"
	"fmt"

	"gitlab.com/Aubichol/hrishi-backend/model"
	storeuser "gitlab.com/Aubichol/hrishi-backend/store/user"
	mongoModel "gitlab.com/Aubichol/hrishi-backend/store/user/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/dig"
)

//users handles user related database queries
type users struct {
	c *mongo.Collection
}

// Save saves user from model to database
func (u *users) Save(modelUser *model.User) error {
	mongoUser := mongoModel.User{}
	if err := mongoUser.FromModel(modelUser); err != nil {
		return fmt.Errorf("could not convert model user to mongo user: %w", err)
	}

	if modelUser.ID == "" {
		mongoUser.ID = primitive.NewObjectID()
	}

	filter := bson.M{"_id": mongoUser.ID}
	update := bson.M{"$set": mongoUser}
	upsert := true

	_, err := u.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)

	return err
}

//SetProfilePic sets the profile picture of the user by giving the pictures "pic" value
func (u *users) SetProfilePic(userID, pic string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"profile_pic": pic,
	}}

	_, err = u.c.UpdateOne(
		context.Background(),
		filter,
		update,
		&options.UpdateOptions{}, //Why update options are not given?
	)

	return err
}

//FindByID finds a user by id
func (u *users) FindByID(id string) (*model.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id %s : %w", id, err)
	}

	filter := bson.M{"_id": objectID}
	result := u.c.FindOne(context.Background(), filter, &options.FindOneOptions{})
	if err := result.Err(); err != nil {
		return nil, err
	}

	user := mongoModel.User{}
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("Could not decode mongo model to model : %w", err)
	}

	return user.ModelUser(), nil
}

//FindByEmail finds a user by email
func (u *users) FindByEmail(email string) (*model.User, error) {
	filter := bson.M{"email": email}
	result := u.c.FindOne(context.Background(), filter, &options.FindOneOptions{})
	if err := result.Err(); err != nil {
		return nil, err
	}

	user := mongoModel.User{}
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("could not decode data from mongo : %w", err)
	}

	return user.ModelUser(), nil
}

//FindByIDs returns all the users from multiple user ids
func (u *users) FindByIDs(ids ...string) ([]*model.User, error) {
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

	cursor, err := u.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return u.cursorToUsers(cursor)
}

//All should return all existing users
func (u *users) All(userID string) ([]*model.User, error) {

	filter := bson.M{"real": true}

	cursor, err := u.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return u.cursorToUsers(cursor)
}

//All should return all existing users
func (u *users) AllPublic() ([]*model.User, error) {

	filter := bson.M{"real": true}

	cursor, err := u.c.Find(context.Background(), filter, nil)
	if err != nil {
		return nil, err
	}

	return u.cursorToUsers(cursor)
}

//Search search for users given the text, skip and limit
func (u *users) Search(text string, skip, limit int64) ([]*model.User, error) {
	filter := bson.M{"$text": bson.M{"$search": text}}
	cursor, err := u.c.Find(context.Background(), filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}

	return u.cursorToUsers(cursor)
}

//cursorToUsers decodes users one by one from the search result
func (u *users) cursorToUsers(cursor *mongo.Cursor) ([]*model.User, error) {
	defer cursor.Close(context.Background())
	modelUsers := []*model.User{}

	for cursor.Next(context.Background()) {
		user := mongoModel.User{}
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("Could not decode data from mongodb %w", err)
		}

		modelUsers = append(modelUsers, user.ModelUser())
	}

	return modelUsers, nil
}

// Params provides parameters for user specific Collection
type Params struct {
	dig.In
	Collection *mongo.Collection `name:"users"`
}

//Store provides store for users
func Store(params Params) storeuser.Users {
	return &users{params.Collection}
}
