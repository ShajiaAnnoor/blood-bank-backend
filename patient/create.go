package patient

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/patient/dto"
	storepatient "gitlab.com/Aubichol/blood-bank-backend/store/patient"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating patient
type Creater interface {
	Create(create *dto.Patient (*dto.CreateResponse, error)
}

// create creates patient
type create struct {
	storePatient storepatient.Patient
	validate    *validator.Validate
}

func (c *create) toModel(userpatient *dto.Patient) (
	patient *model.Patient,
) {
	patient = &model.Patient{}
	patient.CreatedAt = time.Now().UTC()
	patient.UpdatedAt = patient.CreatedAt
	patient.Description = userpatient.Description
	patient.Title = userpatient.Title
	patient.UserID = userpatient.UserID
	return
}

func (c *create) validateData(create *dto.Patient) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Patient) (
	modelPatient *model.Patient,
) {
	modelPatient = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Patient) (
	id string,
	err error,
) {
	id, err = c.storePatient.Save(model)
	return
}

func (c *create) giveResponse(modelPatient *model.Patient, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{
		"id": modelPatient.UserID,
	}).Debug("Patient created successfully")

	return &dto.CreateResponse{
		Message:    "patient created",
		OK:         true,
		PatientTime: modelPatient.CreatedAt.String(),
		ID:         id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create patient. Error: ", err)
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "Invalid data",
		},
	}

	err = fmt.Errorf("%s %w", err.Error(), &errResp)
	return
}

//Create implements Creater interface
func (c *create) Create(create *dto.Patient) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelPatient := c.convertData(create)

	id, err := c.askStore(modelPatient)
	if err == nil {
		return c.giveResponse(modelPatient, id)
	}

	err = c.giveError()
	return nil, err
}

//CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	StorePatients storepatient.Patient
	Validate      *validator.Validate
}

//NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.StorePatients,
		params.Validate,
	}
}
