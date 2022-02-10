package comment

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Comments wraps user's comment functionality
type Comments interface {
	Save(*model.Comment) (id string, err error)
	FindByID(id string) (*model.Comment, error)
	FindByStatusID(id string, skip int64, limit int64) ([]*model.Comment, error)
	CountByStatusID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Comment, error)
	Search(q string, skip, limit int64) ([]*model.Comment, error)
}
