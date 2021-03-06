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

//Deleter provides an interface for updating organizationes
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes organization
type delete struct {
	storeOrganization storeorganization.Organizations
	validate          *validator.Validate
}

func (d *delete) toModel(userorganization *dto.Delete) (organization *model.Organization) {
	organization = &model.Organization{}
	organization.UpdatedAt = time.Now().UTC()
	organization.UserID = userorganization.UserID
	organization.IsDeleted = true
	organization.ID = userorganization.OrganizationID
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelOrganization *model.Organization,
) {
	modelOrganization = d.toModel(delete)
	return
}

func (d *delete) askStore(modelOrganization *model.Organization) (
	id string,
	err error,
) {
	id, err = d.storeOrganization.Save(modelOrganization)
	return
}

func (d *delete) giveResponse(
	modelOrganization *model.Organization,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelOrganization.UserID,
	}).Debug("User deleted organization successfully")

	return &dto.DeleteResponse{
		Message:     "Organization deleted",
		OK:          true,
		ID:          id,
		DeletedTime: modelOrganization.UpdatedAt.String(),
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

	modelOrganization := d.convertData(delete)
	id, err := d.askStore(modelOrganization)
	if err == nil {
		return d.giveResponse(modelOrganization, id), nil
	}

	logrus.Error("Could not delete organization ", err)
	err = d.giveError()
	return nil, err
}

//NewDelete returns new instance of NewDelete
func NewDelete(store storeorganization.Organizations, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
