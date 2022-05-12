package organization

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	storeorganization "gitlab.com/Aubichol/blood-bank-backend/store/organization"
	"go.uber.org/dig"
	validator "gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating user organization
type Creater interface {
	Create(create *dto.Organization) (*dto.CreateResponse, error)
}

// create creates user organization
type create struct {
	storeOrganization storeorganization.Organizations
	validate          *validator.Validate
}

func (c *create) toModel(userorganization *dto.Organization) (
	organization *model.Organization,
) {
	organization = &model.Organization{}
	organization.CreatedAt = time.Now().UTC()
	organization.UpdatedAt = organization.CreatedAt
	organization.Description = userorganization.Description
	organization.Name = userorganization.Name
	organization.Phone = userorganization.Phone
	organization.District = userorganization.District
	organization.Address = userorganization.Address
	organization.UserID = userorganization.UserID
	return
}

func (c *create) validateData(create *dto.Organization) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Organization) (
	modelOrganization *model.Organization,
) {
	modelOrganization = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Organization) (
	id string,
	err error,
) {
	id, err = c.storeOrganization.Save(model)
	return
}

func (c *create) giveResponse(modelOrganization *model.Organization, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{
		"id": modelOrganization.UserID,
	}).Debug("User created organization successfully")

	return &dto.CreateResponse{
		Message: "organization created",
		OK:      true,
		//		OrganizationTime: modelOrganization.CreatedAt.String(),
		ID: id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create organization. Error: ", err)
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
func (c *create) Create(create *dto.Organization) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelOrganization := c.convertData(create)

	id, err := c.askStore(modelOrganization)
	if err == nil {
		return c.giveResponse(modelOrganization, id)
	}

	err = c.giveError()
	return nil, err
}

//CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	StoreOrganizationes storeorganization.Organizations
	Validate            *validator.Validate
}

//NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.StoreOrganizationes,
		params.Validate,
	}
}
