package staticcontent

import "gitlab.com/Aubichol/blood-bank-backend/model"

// StaticContents wraps user's comment functionality
type StaticContents interface {
	Save(*model.StaticContent) (id string, err error)
	FindByID(id string) (*model.StaticContent, error)
	FindByStaticContentID(id string, skip int64, limit int64) ([]*model.StaticContent, error)
	CountByStaticContentID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.StaticContent, error)
	Search(q string, skip, limit int64) ([]*model.StaticContent, error)
}
