package donor

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Donors wraps donor's functionality
type Donors interface {
	Save(*model.Donor) (id string, err error)
	FindByID(id string) (*model.Donor, error)
	FindByDonorID(id string, skip int64, limit int64) ([]*model.Donor, error)
	CountByDonorID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Donor, error)
	Search(q string, skip, limit int64) ([]*model.Donor, error)
}
