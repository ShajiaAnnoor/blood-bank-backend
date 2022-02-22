package staticcontent

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent/dto"
	storestaticcontent "gitlab.com/Aubichol/blood-bank-backend/store/staticcontent"
	"gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for updating staticcontentes
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes staticcontent
type delete struct {
	storeStaticcontent storestaticcontent.Staticcontent
	validate           *validator.Validate
}

func (u *delete) toModel(userstaticcontent *dto.Delete) (staticcontent *model.Staticcontent) {
	staticcontent = &model.Staticcontent{}
	staticcontent.CreatedAt = time.Now().UTC()
	staticcontent.DeletedAt = staticcontent.CreatedAt
	staticcontent.Status = userstaticcontent.Status
	staticcontent.UserID = userstaticcontent.UserID
	staticcontent.ID = userstaticcontent.StatusID
	return
}

func (u *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(u.validate)
	return
}

func (u *delete) convertData(delete *dto.Delete) (
	modelStaticcontent *model.Staticcontent,
) {
	modelStaticcontent = u.toModel(delete)
	return
}

func (u *delete) askStore(modelStaticcontent *model.Staticcontent) (
	id string,
	err error,
) {
	id, err = u.storeStaticcontent.Save(modelStaticcontent)
	return
}

func (u *delete) giveResponse(
	modelStaticcontent *model.Staticcontent,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelStaticcontent.UserID,
	}).Debug("User deleted staticcontent successfully")

	return &dto.DeleteResponse{
		Message:    "Staticcontent deleted",
		OK:         true,
		ID:         id,
		DeleteTime: modelStaticcontent.DeletedAt.String(),
	}
}

func (u *delete) giveError() (err error) {
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
func (u *delete) Delete(delete *dto.Delete) (
	*dto.DeleteResponse, error,
) {
	if err := u.validateData(delete); err != nil {
		return nil, err
	}

	modelStaticcontent := u.convertData(delete)
	id, err := u.askStore(modelStaticcontent)
	if err == nil {
		return u.giveResponse(modelStaticcontent, id), nil
	}

	logrus.Error("Could not delete staticcontent ", err)
	err = u.giveError()
	return nil, err
}

//NewDelete returns new instance of NewDelete
func NewDelete(store storestaticcontent.Staticcontent, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
