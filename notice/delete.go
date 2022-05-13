package notice

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	storenotice "gitlab.com/Aubichol/blood-bank-backend/store/notice"
	"gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for deleting noticees
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes notice
type delete struct {
	storeNotice storenotice.Notices
	validate    *validator.Validate
}

func (d *delete) toModel(usernotice *dto.Delete) (notice *model.Notice) {
	notice = &model.Notice{}
	notice.UpdatedAt = time.Now().UTC()
	notice.UserID = usernotice.UserID
	notice.ID = usernotice.NoticeID
	notice.IsDeleted = true
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelNotice *model.Notice,
) {
	modelNotice = d.toModel(delete)
	return
}

func (d *delete) askStore(modelNotice *model.Notice) (
	id string,
	err error,
) {
	id, err = d.storeNotice.Save(modelNotice)
	return
}

func (d *delete) giveResponse(
	modelNotice *model.Notice,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelNotice.UserID,
	}).Debug("User deleted notice successfully")

	return &dto.DeleteResponse{
		Message:     "Notice deleted",
		OK:          true,
		ID:          id,
		DeletedTime: modelNotice.UpdatedAt.String(),
	}
}

func (d *delete) giveError() (err error) {
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "Invalid data",
		},
	}
	err = fmt.Errorf(
		"%s %w",
		err.Error(),
		&errResp,
	)
	return
}

//Delete implements Delete interface
func (d *delete) Delete(delete *dto.Delete) (
	*dto.DeleteResponse, error,
) {
	if err := d.validateData(delete); err != nil {
		return nil, err
	}

	modelNotice := d.convertData(delete)
	id, err := d.askStore(modelNotice)
	if err == nil {
		return d.giveResponse(modelNotice, id), nil
	}

	logrus.Error("Could not delete notice ", err)
	err = d.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storenotice.Notices, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
