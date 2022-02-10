package status

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/status/dto"
	storenotice "gitlab.com/Aubichol/blood-bank-backend/store/notice"
	storestatus "gitlab.com/Aubichol/blood-bank-backend/store/status"
	"gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating statuses
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates user status
type update struct {
	storeStatus storenotice.Notice
	validate    *validator.Validate
}

func (u *update) toModel(userstatus *dto.Update) (status *model.Status) {
	status = &model.Status{}
	status.CreatedAt = time.Now().UTC()
	status.UpdatedAt = status.CreatedAt
	status.Status = userstatus.Status
	status.UserID = userstatus.UserID
	status.ID = userstatus.StatusID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelStatus *model.Status,
) {
	modelStatus = u.toModel(update)
	return
}

func (u *update) askStore(modelStatus *model.Status) (
	id string,
	err error,
) {
	id, err = u.storeStatus.Save(modelStatus)
	return
}

func (u *update) giveResponse(
	modelStatus *model.Status,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelStatus.UserID,
	}).Debug("User updated status successfully")

	return &dto.UpdateResponse{
		Message:    "Status updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelStatus.UpdatedAt.String(),
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
func (u *update) Update(update *dto.Update) (
	*dto.UpdateResponse, error,
) {
	if err := u.validateData(update); err != nil {
		return nil, err
	}

	modelStatus := u.convertData(update)
	id, err := u.askStore(modelStatus)
	if err == nil {
		return u.giveResponse(modelStatus, id), nil
	}

	logrus.Error("Could not update status ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewCreate
func NewUpdate(store storestatus.Status, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
