package donor

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/donor/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storedonor "gitlab.com/Aubichol/blood-bank-backend/store/donor"
	validator "gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for updating noticees
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes notice
type delete struct {
	storeDonor storedonor.Donors
	validate   *validator.Validate
}

func (d *delete) toModel(userdonor *dto.Delete) (donor *model.Donor) {
	donor = &model.Donor{}

	donor.UpdatedAt = time.Now().UTC()
	donor.IsDeleted = true
	donor.UserID = userdonor.UserID
	donor.ID = userdonor.DonorID
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelDonor *model.Donor,
) {
	modelDonor = d.toModel(delete)
	return
}

func (d *delete) askStore(modelDonor *model.Donor) (
	id string,
	err error,
) {
	id, err = d.storeDonor.Save(modelDonor)
	return
}

func (d *delete) giveResponse(
	modelNotice *model.Donor,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelNotice.UserID,
	}).Debug("User deleted donor successfully")

	return &dto.DeleteResponse{
		Message: "Donor deleted",
		OK:      true,
		ID:      id,
		//		DeleteTime: modelNotice.DeletedAt.String(),
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

	modelDonor := d.convertData(delete)
	id, err := d.askStore(modelDonor)
	if err == nil {
		return d.giveResponse(modelDonor, id), nil
	}

	logrus.Error("Could not delete notice ", err)
	err = d.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storedonor.Donors, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
