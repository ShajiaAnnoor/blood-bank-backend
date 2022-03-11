package bloodrequest

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storebloodreq "gitlab.com/Aubichol/blood-bank-backend/store/bloodrequest"
	"gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating comment
type Creater interface {
	Create(create *dto.BloodRequest) (*dto.CreateResponse, error)
}

// create creates user blood requests
type create struct {
	storeBloodRequest storebloodreq.BloodRequests
	validate          *validator.Validate
}

func (c *create) toModel(userbloodreq *dto.BloodReq) (bloodreq *model.BloodRequest) {
	bloodreq = &model.BloodRequest{}
	bloodreq.CreatedAt = time.Now().UTC()
	bloodreq.UpdatedAt = bloodreq.CreatedAt
	//	bloodreq.Comment = userbloodreq.Request
	bloodreq.UserID = userbloodreq.UserID
	//	bloodreq.StatusID = userbloodreq.StatusID
	return
}

func (c *create) validateData(create *dto.BloodReq) (err error) {
	err = create.Validate(c.validate)
	return err
}

func (c *create) convertData(create *dto.BloodReq) (
	modelBloodReq *model.BloodRequest,
) {
	modelBloodReq = c.toModel(create)
	return
}

func (c *create) askStore(modelBloodReq *model.BloodRequest) (
	id string,
	err error,
) {
	id, err = c.storeBloodRequest.Save(modelBloodReq)
	return
}

func (c *create) printLog(id string) {
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Debug("User created blood request successfully")
}

func (c *create) createResponse(
	modelBloodReq *model.BloodRequest,
	id string,
) (resp *dto.CreateResponse) {
	return &dto.CreateResponse{
		Message: "blood request created",
		OK:      true,
		ID:      id,
		//		BloodReqTime: modelBloodReq.CreatedAt.String(),
	}
}

func (c *create) giveError(err error) error {
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "invalid data",
		},
	}
	err = fmt.Errorf("%s %w", err.Error(), &errResp)
	return err
}

func (c *create) responseError(err error) (
	*dto.CreateResponse, error) {
	return nil, err
}

func (c *create) giveResponse(resp *dto.CreateResponse) (
	*dto.CreateResponse,
	error,
) {
	return resp, nil
}

func (c *create) noError(err error) (noerror bool) {
	if err == nil {
		noerror = true
	}
	return noerror
}

func (c *create) logError(message string, err error) {
	logrus.Error(message, err)
}

//Create implements Creater interface
func (c *create) Create(create *dto.BloodReq) (*dto.CreateResponse, error) {
	if err := c.validateData(create); err != nil {
		return c.responseError(err)
	}

	modelBloodReq := c.convertData(create)

	id, err := c.askStore(modelBloodReq)
	if c.noError(err) {
		c.printLog(modelBloodReq.ID)
		return c.giveResponse(
			c.createResponse(modelBloodReq, id),
		)
	}
	message := "Could not create blood request "
	c.logError(message, err)
	return c.responseError(c.giveError(err))
}

//NewCreate returns new instance of Creater
func NewCreate(store storebloodreq.BloodRequests, validate *validator.Validate) Creater {
	return &create{
		store,
		validate,
	}
}
