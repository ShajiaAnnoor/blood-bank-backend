package index

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"go.uber.org/dig"
)

func Create(c container.Container) {
	createUsers(c)
	createFriendRequests(c)
	createConversations(c)
	createConversationReadPointer(c)
}

func createUsers(c container.Container) {
	c.Resolve(func(params struct {
		dig.In

		Collection *mongo.Collection `name:"users"`
	}) {

		unique := true
		defaultLanguage := "none"

		models := []mongo.IndexModel{
			{
				Keys: bsonx.Doc{
					{Key: "first_name", Value: bsonx.String("text")},
					{Key: "last_name", Value: bsonx.String("text")},
					{Key: "email", Value: bsonx.String("text")},
				},
				Options: &options.IndexOptions{
					Name:            &usersIndexSearch,
					DefaultLanguage: &defaultLanguage,
				},
			},
			{
				Keys: bsonx.Doc{
					{Key: "email", Value: bsonx.Int32(1)},
				},
				Options: &options.IndexOptions{
					Unique: &unique,
					Name:   &usersIndexEmail,
				},
			},
		}

		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		_, err := params.Collection.Indexes().CreateMany(context.Background(), models, opts)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("index created for users")
	})
}

func createFriendRequests(c container.Container) {
	c.Resolve(func(params struct {
		dig.In

		Collection *mongo.Collection `name:"friend_requests"`
	}) {

		unique := true

		models := []mongo.IndexModel{
			{
				Keys: bsonx.Doc{
					{Key: "from_user_id", Value: bsonx.Int32(1)},
				},
				Options: &options.IndexOptions{Name: &friendRequestIndexFromUserID},
			},
			{
				Keys: bsonx.Doc{
					{Key: "to_user_id", Value: bsonx.Int32(1)},
				},
				Options: &options.IndexOptions{Name: &friendRequestIndexToUserID},
			},
			{
				Keys: bsonx.Doc{
					{Key: "unique_tag", Value: bsonx.Int32(1)},
				},
				Options: &options.IndexOptions{
					Name:   &friendRequestIndexUniqueTag,
					Unique: &unique,
				},
			},
		}

		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		_, err := params.Collection.Indexes().CreateMany(context.Background(), models, opts)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("index created for friend_requests")
	})
}

func createConversations(c container.Container) {
	c.Resolve(func(params struct {
		dig.In

		Collection *mongo.Collection `name:"conversations"`
	}) {

		models := []mongo.IndexModel{
			{
				Keys: bsonx.Doc{
					{Key: "unique_tag", Value: bsonx.Int32(1)},
				},
				Options: &options.IndexOptions{Name: &conversationsUniqueTag},
			},
		}

		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		_, err := params.Collection.Indexes().CreateMany(context.Background(), models, opts)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("index created for conversations")
	})
}

func createConversationReadPointer(c container.Container) {
	c.Resolve(func(params struct {
		dig.In

		Collection *mongo.Collection `name:"conversation_read_pointers"`
	}) {

		unique := true
		models := []mongo.IndexModel{
			{
				Keys: bsonx.Doc{
					{Key: "unique_tag", Value: bsonx.Int32(1)},
					{Key: "user_id", Value: bsonx.Int32(1)},
				},
				Options: &options.IndexOptions{
					Name:   &conversationReadPointersUniqueTagUserID,
					Unique: &unique,
				},
			},
		}

		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		_, err := params.Collection.Indexes().CreateMany(context.Background(), models, opts)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("index created for conversation_read_pointers")
	})
}
