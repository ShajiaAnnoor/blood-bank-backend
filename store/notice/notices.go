package notice

import "gitlab.com/Aubichol/blood-bank-backend/model"

// Comments wraps user's comment functionality
type Notices interface {
	Save(*model.Notice) (id string, err error)
	FindByID(id string) (*model.Notice, error)
	FindByNoticeID(id string, skip int64, limit int64) ([]*model.Notice, error)
	CountByNoticeID(id string) (int64, error)
	FindByIDs(id ...string) ([]*model.Notice, error)
	Search(q string, skip, limit int64) ([]*model.Notice, error)
}
