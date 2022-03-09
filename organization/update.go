package organization

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	storeorganization "gitlab.com/Aubichol/blood-bank-backend/store/organization"
	validator "gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating organizationes
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates user organization
type update struct {
	storeorganization storeorganization.Organization
	validate          *validator.Validate
}

func (u *update) toModel(userorganization *dto.Update) (organization *model.Organization) {
	organization = &model.Organization{}
	organization.CreatedAt = time.Now().UTC()
	organization.UpdatedAt = organization.CreatedAt
	organization.Organization = userorganization.Organization
	organization.UserID = userorganization.UserID
	organization.ID = userorganization.OrganizationID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelOrganization *model.Organization,
) {
	modelOrganization = u.toModel(update)
	return
}

func (u *update) askStore(modelOrganization *model.Organization) (
	id string,
	err error,
) {
	id, err = u.storeOrganization.Save(modelOrganization)
	return
}

func (u *update) giveResponse(
	modelOrganization *model.Organization,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelOrganization.UserID,
	}).Debug("User updated organization successfully")

	return &dto.UpdateResponse{
		Message:    "Organization updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelOrganization.UpdatedAt.String(),
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

	modelOrganization := u.convertData(update)
	id, err := u.askStore(modelOrganization)
	if err == nil {
		return u.giveResponse(modelOrganization, id), nil
	}

	logrus.Error("Could not update organization ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of update
func NewUpdate(store storeorganization.Organization, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
