package user

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Users wraps user's store functionality
type Users interface {
	Save(*model.User) error
	SetProfilePic(userID, pic string) error
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByIDs(id ...string) ([]*model.User, error)
	All(userID string) ([]*model.User, error)
	AllPublic() ([]*model.User, error)
	Search(q string, skip, limit int64) ([]*model.User, error)
}
