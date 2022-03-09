package donor

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/donor/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storedonor "gitlab.com/Aubichol/blood-bank-backend/store/donor"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating donor
type Creater interface {
	Create(create *dto.Donor) (*dto.CreateResponse, error)
}

// create creates donor
type create struct {
	storeDonor storedonor.Donor
	validate   *validator.Validate
}

func (c *create) toModel(userdonor *dto.Donor) (
	donor *model.Donor,
) {
	donor = &model.Donor{}
	donor.CreatedAt = time.Now().UTC()
	donor.UpdatedAt = donor.CreatedAt
	donor.Description = userdonor.Description
	donor.Title = userdonor.Title
	donor.UserID = userdonor.UserID
	return
}

func (c *create) validateData(create *dto.Donor) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Donor) (
	modelDonor *model.Donor,
) {
	modelDonor = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Donor) (
	id string,
	err error,
) {
	id, err = c.storeDonor.Save(model)
	return
}

func (c *create) giveResponse(modelDonor *model.Donor, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{
		"id": modelDonor.UserID,
	}).Debug("User created donor successfully")

	return &dto.CreateResponse{
		Message:   "donor created",
		OK:        true,
		DonorTime: modelDonor.CreatedAt.String(),
		ID:        id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create donor. Error: ", err)
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
func (c *create) Create(create *dto.Donor) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelDonor := c.convertData(create)

	id, err := c.askStore(modelDonor)
	if err == nil {
		return c.giveResponse(modelDonor, id)
	}

	err = c.giveError()
	return nil, err
}

//CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	StoreDonors storedonor.Donor
	Validate    *validator.Validate
}

//NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.StoreDonors,
		params.Validate,
	}
}
