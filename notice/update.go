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

//Updater provides an interface for updating noticees
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates user notice
type update struct {
	storeNotice storenotice.Notice
	validate    *validator.Validate
}

func (u *update) toModel(usernotice *dto.Update) (notice *model.Notice) {
	notice = &model.Notice{}
	notice.CreatedAt = time.Now().UTC()
	notice.UpdatedAt = notice.CreatedAt
	notice.Status = usernotice.Status
	notice.UserID = usernotice.UserID
	notice.ID = usernotice.StatusID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(upd *dto.Update) (
	modelNotice *model.Notice,
) {
	modelNotice = u.toModel(upd)
	return
}

func (u *update) askStore(modelNotice *model.Notice) (
	id string,
	err error,
) {
	id, err = u.storeNotice.Save(modelNotice)
	return
}

func (u *update) giveResponse(
	modelNotice *model.Notice,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelNotice.UserID,
	}).Debug("Updated notice successfully")

	return &dto.UpdateResponse{
		Message:    "Notice updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelNotice.UpdatedAt.String(),
	}
}

func (u *update) giveError() (err error) {
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

//Update implements Update interface
func (u *update) Update(upd *dto.Update) (
	*dto.UpdateResponse, error,
) {
	if err := u.validateData(upd); err != nil {
		return nil, err
	}

	modelNotice := u.convertData(upd)
	id, err := u.askStore(modelNotice)
	if err == nil {
		return u.giveResponse(modelNotice, id), nil
	}

	logrus.Error("Could not update notice ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewUpdate
func NewUpdate(store storenotice.Notice, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
