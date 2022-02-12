package bloodrequest

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank/model"
	storebloodreq "gitlab.com/Aubichol/blood-bank/store/bloodrequest"
	"gopkg.in/go-playground/validator.v9"
)

// Creater provides create method for creating comment
type Creater interface {
	Create(create *dto.Bloodreq) (*dto.CreateResponse, error)
}

// create creates user comment
type create struct {
	storeBloodRequest storebloodreq.BloodRequests
	validate          *validator.Validate
}

func (c *create) toModel(userbloodreq *dto.Comment) (bloodreq *model.Comment) {
	bloodreq = &model.Bloodreq{}
	bloodreq.CreatedAt = time.Now().UTC()
	bloodreq.UpdatedAt = bloodreq.CreatedAt
	bloodreq.Comment = userbloodreq.Comment
	bloodreq.UserID = userbloodreq.UserID
	bloodreq.StatusID = userbloodreq.StatusID
	return
}

func (c *create) validateData(create *dto.Comment) (err error) {
	err = create.Validate(c.validate)
	return err
}

func (c *create) convertData(create *dto.Comment) (
	modelComment *model.Comment,
) {
	modelComment = c.toModel(create)
	return
}

func (c *create) askStore(modelComment *model.Comment) (
	id string,
	err error,
) {
	id, err = c.storeComment.Save(modelComment)
	return
}

func (c *create) printLog(id string) {
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Debug("User created comment successfully")
}

func (c *create) createResponse(
	modelComment *model.Comment,
	id string,
) (resp *dto.CreateResponse) {
	return &dto.CreateResponse{
		Message:     "comment created",
		OK:          true,
		ID:          id,
		CommentTime: modelComment.CreatedAt.String(),
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

func (c *create) resopnseError(err error) (
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
func (c *create) Create(create *dto.Comment) (*dto.CreateResponse, error) {
	if err := c.validateData(create); err != nil {
		return c.resopnseError(err)
	}

	modelComment := c.convertData(create)

	id, err := c.askStore(modelComment)
	if c.noError(err) {
		c.printLog(modelComment.ID)
		return c.giveResponse(
			c.createResponse(modelComment, id),
		)
	}
	message := "Could not create comment "
	c.logError(message, err)
	return c.resopnseError(c.giveError(err))
}

//NewCreate returns new instance of Creater
func NewCreate(store storecomment.Comments, validate *validator.Validate) Creater {
	return &create{
		store,
		validate,
	}
}
