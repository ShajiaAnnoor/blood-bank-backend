package bloodrequest

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storebloodrequest "gitlab.com/Aubichol/blood-bank-backend/store/bloodrequest"
	"gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for updating bloodrequestes
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes bloodrequest
type delete struct {
	storeBloodRequest storebloodrequest.BloodRequest
	validate          *validator.Validate
}

func (u *delete) toModel(userbloodrequest *dto.Delete) (bloodrequest *model.BloodRequest) {
	bloodrequest = &model.BloodRequest{}
	bloodrequest.CreatedAt = time.Now().UTC()
	bloodrequest.DeletedAt = bloodrequest.CreatedAt
	bloodrequest.Status = userbloodrequest.Status
	bloodrequest.UserID = userbloodrequest.UserID
	bloodrequest.ID = userbloodrequest.StatusID
	return
}

func (u *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(u.validate)
	return
}

func (u *delete) convertData(delete *dto.Delete) (
	modelBloodRequest *model.BloodRequest,
) {
	modelBloodRequest = u.toModel(delete)
	return
}

func (u *delete) askStore(modelBloodRequest *model.BloodRequest) (
	id string,
	err error,
) {
	id, err = u.storeBloodRequest.Save(modelBloodRequest)
	return
}

func (u *delete) giveResponse(
	modelBloodRequest *model.BloodRequest,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelBloodRequest.UserID,
	}).Debug("User deleted bloodrequest successfully")

	return &dto.DeleteResponse{
		Message:    "BloodRequest deleted",
		OK:         true,
		ID:         id,
		DeleteTime: modelBloodRequest.DeletedAt.String(),
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

	modelBloodRequest := u.convertData(delete)
	id, err := u.askStore(modelBloodRequest)
	if err == nil {
		return u.giveResponse(modelBloodRequest, id), nil
	}

	logrus.Error("Could not delete bloodrequest ", err)
	err = u.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storebloodrequest.BloodRequest, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
