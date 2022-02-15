package staticcontent

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent/dto"
	storestaticcontent "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating user staticcontent
type Creater interface {
	Create(create *dto.Staticcontent) (*dto.CreateResponse, error)
}

// create creates user staticcontent
type create struct {
	storeStaticcontent storestaticcontent.Staticcontent
	validate           *validator.Validate
}

func (c *create) toModel(userstaticcontent *dto.StaticContent) (
	staticcontent *model.StaticContent,
) {
	sc = &model.StaticContent{}
	sc.CreatedAt = time.Now().UTC()
	sc.UpdatedAt = staticcontent.CreatedAt
	sc.Description = userstaticcontent.Description
	sc.Title = userstaticcontent.Title
	sc.UserID = userstaticcontent.UserID
	return
}

func (c *create) validateData(create *dto.Staticcontent) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Staticcontent) (
	modelStaticcontent *model.Staticcontent,
) {
	modelStaticcontent = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Staticcontent) (
	id string,
	err error,
) {
	id, err = c.storeStaticcontent.Save(model)
	return
}

func (c *create) giveResponse(modelStaticcontent *model.Staticcontent, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{
		"id": modelStaticcontent.UserID,
	}).Debug("User created staticcontent successfully")

	return &dto.CreateResponse{
		Message:           "staticcontent created",
		OK:                true,
		StaticcontentTime: modelStaticcontent.CreatedAt.String(),
		ID:                id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create staticcontent. Error: ", err)
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
func (c *create) Create(create *dto.Staticcontent) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelStaticcontent := c.convertData(create)

	id, err := c.askStore(modelStaticcontent)
	if err == nil {
		return c.giveResponse(modelStaticcontent, id)
	}

	err = c.giveError()
	return nil, err
}

//CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	StoreStaticcontentes storestaticcontent.Staticcontent
	Validate             *validator.Validate
}

//NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.StoreStaticcontentes,
		params.Validate,
	}
}
