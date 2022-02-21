package donor

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/donor/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storedonor "gitlab.com/Aubichol/blood-bank-backend/store/donor"
	"gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating donors
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates donor
type update struct {
	storeDonor storedonor.Donor
	validate   *validator.Validate
}

func (u *update) toModel(userdonor *dto.Update) (donor *model.Donor) {
	donor = &model.Donor{}
	donor.CreatedAt = time.Now().UTC()
	donor.UpdatedAt = donor.CreatedAt
	donor.UserID = userdonor.UserID
	donor.ID = userdonor.StatusID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelDonor *model.Donor,
) {
	modelDonor = u.toModel(update)
	return
}

func (u *update) askStore(modelStatus *model.Donor) (
	id string,
	err error,
) {
	id, err = u.storeDonor.Save(modelDonor)
	return
}

func (u *update) giveResponse(
	modelDonor *model.Donor,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelDonor.UserID,
	}).Debug("User updated donor successfully")

	return &dto.UpdateResponse{
		Message:    "Donor updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelDonor.UpdatedAt.String(),
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

	modelDonor := u.convertData(update)
	id, err := u.askStore(modelDonor)
	if err == nil {
		return u.giveResponse(modelDonor, id), nil
	}

	logrus.Error("Could not update donor ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewCreate
func NewUpdate(store storedonor.Donor, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
