package blood-request

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Comments wraps user's blood request functionality
type BloodRequests interface {
	Save(*model.BloodRequest) (id string, err error)
	FindByID(id string) (*model.BloodRequest, error)
	FindByBloodRequestID(id string, skip int64, limit int64) ([]*model.BloodRequest, error)
	CountByBloodRequestID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.BloodRequest, error)
	Search(q string, skip, limit int64) ([]*model.BloodRequest, error)
}
