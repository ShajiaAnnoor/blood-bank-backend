package organization

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Comments wraps user's comment functionality
type Organizations interface {
	Save(*model.Organization) (id string, err error)
	FindByID(id string) (*model.Organization, error)
	FindByOrganizationID(id string, skip int64, limit int64) ([]*model.Organization, error)
	CountByOrganizationID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Organization, error)
	Search(q string, skip, limit int64) ([]*model.Organization, error)
}
