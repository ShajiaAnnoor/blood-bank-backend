package patient

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/patient/dto"
	storepatient "gitlab.com/Aubichol/blood-bank-backend/store/patient"
	"gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for updating patientes
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes patient
type delete struct {
	storePatient storepatient.Patient
	validate     *validator.Validate
}

func (u *delete) toModel(userpatient *dto.Delete) (patient *model.Patient) {
	patient = &model.Patient{}
	patient.CreatedAt = time.Now().UTC()
	patient.DeletedAt = patient.CreatedAt
	patient.Status = userpatient.Status
	patient.UserID = userpatient.UserID
	patient.ID = userpatient.StatusID
	return
}

func (u *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(u.validate)
	return
}

func (u *delete) convertData(delete *dto.Delete) (
	modelPatient *model.Patient,
) {
	modelPatient = u.toModel(delete)
	return
}

func (u *delete) askStore(modelPatient *model.Patient) (
	id string,
	err error,
) {
	id, err = u.storePatient.Save(modelPatient)
	return
}

func (u *delete) giveResponse(
	modelPatient *model.Patient,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelPatient.UserID,
	}).Debug("User deleted patient successfully")

	return &dto.DeleteResponse{
		Message:    "Patient deleted",
		OK:         true,
		ID:         id,
		DeleteTime: modelPatient.DeletedAt.String(),
	}
}

func (u *delete) giveError() (err error) {
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
func (u *delete) Delete(delete *dto.Delete) (
	*dto.DeleteResponse, error,
) {
	if err := u.validateData(delete); err != nil {
		return nil, err
	}

	modelPatient := u.convertData(delete)
	id, err := u.askStore(modelPatient)
	if err == nil {
		return u.giveResponse(modelPatient, id), nil
	}

	logrus.Error("Could not delete patient ", err)
	err = u.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storepatient.Patient, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
