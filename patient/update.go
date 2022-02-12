package patient

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/status/dto"
	storestatus "gitlab.com/Aubichol/blood-bank-backend/store/status"
	"gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating statuses
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates user status
type update struct {
	storeStatus storestatus.Status
	validate    *validator.Validate
}

func (u *update) toModel(userpatient *dto.Update) (patient *model.Patient) {
	patient = &model.Status{}
	patient.CreatedAt = time.Now().UTC()
	patient.UpdatedAt = patient.CreatedAt
	patient.Status = userpatient.Status
	patient.UserID = userpatient.UserID
	patient.ID = userpatient.StatusID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelPatient *model.Status,
) {
	modelPatient = u.toModel(update)
	return
}

func (u *update) askStore(modelPatient *model.Patient) (
	id string,
	err error,
) {
	id, err = u.storePatient.Save(modelPatient)
	return
}

func (u *update) giveResponse(
	modelPatient *model.Patient,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelPatient.UserID,
	}).Debug("User updated status successfully")

	return &dto.UpdateResponse{
		Message:    "Status updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelPatient.UpdatedAt.String(),
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

	modelPatient := u.convertData(update)
	id, err := u.askStore(modelPatient)
	if err == nil {
		return u.giveResponse(modelPatient, id), nil
	}

	logrus.Error("Could not update patient ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewCreate
func NewUpdate(store storepatient.Patient, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
