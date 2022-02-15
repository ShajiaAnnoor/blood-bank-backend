package comment

import "gitlab.com/Aubichol/hrishi-backend/model"

// Patients wraps patients functionality
type Patients interface {
	Save(*model.Patient) (id string, err error)
	FindByID(id string) (*model.Patient, error)
	FindByPatientID(id string, skip int64, limit int64) ([]*model.Patient, error)
	CountByPatientID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Patient, error)
	Search(q string, skip, limit int64) ([]*model.Patient, error)
}
