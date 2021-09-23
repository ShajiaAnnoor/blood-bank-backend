package status

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/errors"
	"gitlab.com/Aubichol/hrishi-backend/model"
	"gitlab.com/Aubichol/hrishi-backend/status/dto"
	storestatus "gitlab.com/Aubichol/hrishi-backend/store/status"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating user status
type Creater interface {
	Create(create *dto.Status) (*dto.CreateResponse, error)
}

// create creates user status
type create struct {
	storeStatus storestatus.Status
	validate    *validator.Validate
}

func (c *create) toModel(userstatus *dto.Status) (
	status *model.Status,
) {
	status = &model.Status{}
	status.CreatedAt = time.Now().UTC()
	status.UpdatedAt = status.CreatedAt
	status.Status = userstatus.Status
	status.UserID = userstatus.UserID
	return
}

func (c *create) validateData(create *dto.Status) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Status) (
	modelStatus *model.Status,
) {
	modelStatus = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Status) (
	id string,
	err error,
) {
	id, err = c.storeStatus.Save(model)
	return
}

func (c *create) giveResponse(modelStatus *model.Status, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{
		"id": modelStatus.UserID,
	}).Debug("User created status successfully")

	return &dto.CreateResponse{
		Message:    "status created",
		OK:         true,
		StatusTime: modelStatus.CreatedAt.String(),
		ID:         id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create status. Error: ", err)
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
func (c *create) Create(create *dto.Status) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelStatus := c.convertData(create)

	id, err := c.askStore(modelStatus)
	if err == nil {
		return c.giveResponse(modelStatus, id)
	}

	err = c.giveError()
	return nil, err
}

//CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	StoreStatuses storestatus.Status
	Validate      *validator.Validate
}

//NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.StoreStatuses,
		params.Validate,
	}
}
