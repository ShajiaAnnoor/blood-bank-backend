package notice

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	storenotice "gitlab.com/Aubichol/blood-bank-backend/store/notice"
	"go.uber.org/dig"
	"gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating user notice
type Creater interface {
	Create(create *dto.Notice) (*dto.CreateResponse, error)
}

// create creates user notice
type create struct {
	storeNotice storenotice.Notice
	validate    *validator.Validate
}

func (c *create) toModel(usernotice *dto.Notice) (
	notice *model.Notice,
) {
	notice = &model.Notice{}
	notice.CreatedAt = time.Now().UTC()
	notice.UpdatedAt = notice.CreatedAt
	notice.Description = usernotice.Description
	notice.Title = usernotice.Title
	notice.UserID = usernotice.UserID
	return
}

func (c *create) validateData(create *dto.Notice) (
	err error,
) {
	err = create.Validate(c.validate)
	return
}

func (c *create) convertData(create *dto.Notice) (
	modelNotice *model.Notice,
) {
	modelNotice = c.toModel(create)
	return
}

func (c *create) askStore(model *model.Notice) (
	id string,
	err error,
) {
	id, err = c.storeNotice.Save(model)
	return
}

func (c *create) giveResponse(modelNotice *model.Notice, id string) (
	*dto.CreateResponse, error,
) {
	logrus.WithFields(logrus.Fields{
		"id": modelNotice.UserID,
	}).Debug("User created notice successfully")

	return &dto.CreateResponse{
		Message:    "notice created",
		OK:         true,
		NoticeTime: modelNotice.CreatedAt.String(),
		ID:         id,
	}, nil
}

func (c *create) giveError() (err error) {
	logrus.Error("Could not create notice. Error: ", err)
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
func (c *create) Create(create *dto.Notice) (
	*dto.CreateResponse, error,
) {
	err := c.validateData(create)
	if err != nil {
		return nil, err
	}

	modelNotice := c.convertData(create)

	id, err := c.askStore(modelNotice)
	if err == nil {
		return c.giveResponse(modelNotice, id)
	}

	err = c.giveError()
	return nil, err
}

//CreateParams give parameters for NewCreate
type CreateParams struct {
	dig.In
	StoreNotices storenotice.Notice
	Validate     *validator.Validate
}

//NewCreate returns new instance of NewCreate
func NewCreate(params CreateParams) Creater {
	return &create{
		params.StoreNotices,
		params.Validate,
	}
}
