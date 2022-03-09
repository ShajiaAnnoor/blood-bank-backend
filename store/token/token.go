package token

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Token wraps token's store functionality
type Token interface {
	Save(*model.Token) error
	FindByID(id string) (*model.Token, error)
	FindByUser(id string) ([]*model.Token, error)
	FindByIDs(id ...string) ([]*model.Token, error)
	Search(q string, skip, limit int64) ([]*model.Token, error)
}
