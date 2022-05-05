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
	storeBloodRequest storebloodrequest.BloodRequests
	validate          *validator.Validate
}

func (d *delete) toModel(userbloodrequest *dto.Delete) (bloodrequest *model.BloodRequest) {
	bloodrequest = &model.BloodRequest{}
	bloodrequest.UpdatedAt = time.Now().UTC()
	bloodrequest.IsDeleted = true
	//	bloodrequest.DeletedAt = userbloodrequest.DeletedAt
	//	bloodrequest.UserID = userbloodrequest.UserID
	bloodrequest.ID = userbloodrequest.ID
	return
}

func (d *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(d.validate)
	return
}

func (d *delete) convertData(delete *dto.Delete) (
	modelBloodRequest *model.BloodRequest,
) {
	modelBloodRequest = d.toModel(delete)
	return
}

func (d *delete) askStore(modelBloodRequest *model.BloodRequest) (
	id string,
	err error,
) {
	id, err = d.storeBloodRequest.Save(modelBloodRequest)
	return
}

func (d *delete) giveResponse(
	modelBloodRequest *model.BloodRequest,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelBloodRequest.UserID,
	}).Debug("User deleted bloodrequest successfully")

	return &dto.DeleteResponse{
		Message: "BloodRequest deleted",
		OK:      true,
		ID:      id,
		//		DeleteTime: modelBloodRequest.DeletedAt.String(),
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

	modelBloodRequest := d.convertData(delete)
	id, err := d.askStore(modelBloodRequest)
	if err == nil {
		return d.giveResponse(modelBloodRequest, id), nil
	}

	logrus.Error("Could not delete bloodrequest ", err)
	err = d.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storebloodrequest.BloodRequests, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
