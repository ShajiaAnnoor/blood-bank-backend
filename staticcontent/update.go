package staticcontent

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent/dto"
	"gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating staticcontentes
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates user staticcontent
type update struct {
	storeStatus storestaticcontent.StaticContent
	validate    *validator.Validate
}

func (u *update) toModel(userstaticcontent *dto.Update) (sc *model.StaticContent) {
	sc = &model.StaticContent{}
	sc.CreatedAt = time.Now().UTC()
	sc.UpdatedAt = sc.CreatedAt
	sc.Status = userstaticcontent.Status
	sc.UserID = userstaticcontent.UserID
	sc.ID = userstaticcontent.StatusID
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
	}).Debug("User updated staticcontent successfully")

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

	modelStaticContent := u.convertData(update)
	id, err := u.askStore(modelStaticContent)
	if err == nil {
		return u.giveResponse(modelStaticContent, id), nil
	}

	logrus.Error("Could not update staticcontent ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewCreate
func NewUpdate(store storestaticcontent.StaticContent, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
