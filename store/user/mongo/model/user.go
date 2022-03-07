package model

import (
	"time"

	"gitlab.com/Aubichol/hrishi-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User defines mongodb data type for User
type User struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty"`
	FirstName  string                 `bson:"first_name"`
	LastName   string                 `bson:"last_name"`
	Gender     string                 `bson:"gender"`
	BirthDate  BirthDate              `bson:"birth_date"`
	Email      string                 `bson:"email"`
	Password   string                 `bson:"password"`
	Verified   bool                   `bson:"verified"`
	Profile    map[string]interface{} `bson:"profile"`
	CreatedAt  time.Time              `bson:"created_at"`
	UpdatedAt  time.Time              `bson:"updated_at"`
	Real       bool                   `bson:"real"`
	ProfilePic string                 `bson:"profile_pic"`
}

//FromModel converts model data to mongodb model data for user
func (u *User) FromModel(modelUser *model.User) error {
	u.FirstName = modelUser.FirstName
	u.LastName = modelUser.LastName
	u.Gender = modelUser.Gender
	u.BirthDate.FromModel(&modelUser.BirthDate)
	u.Email = modelUser.Email
	u.Password = modelUser.Password
	u.Verified = modelUser.Verified
	u.CreatedAt = modelUser.CreatedAt
	u.UpdatedAt = modelUser.UpdatedAt
	u.Profile = modelUser.Profile
	u.ProfilePic = modelUser.ProfilePic

	if modelUser.ID == "" {
		return nil
	}

	id, err := primitive.ObjectIDFromHex(modelUser.ID)
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

//ModelUser converts bson to model for user
func (u *User) ModelUser() *model.User {
	user := model.User{}
	user.ID = u.ID.Hex()
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Gender = u.Gender
	user.BirthDate = *u.BirthDate.ModelBirthDate()
	user.Email = u.Email
	user.Password = u.Password
	user.Verified = u.Verified
	user.CreatedAt = u.CreatedAt
	user.UpdatedAt = u.UpdatedAt
	user.Profile = u.Profile
	user.ProfilePic = u.ProfilePic

	return &user
}
