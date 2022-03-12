package bloodrequest

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storebloodrequest "gitlab.com/Aubichol/blood-bank-backend/store/bloodrequest"
	validator "gopkg.in/go-playground/validator.v9"
)

//Updater provides an interface for updating bloodrequests
type Updater interface {
	Update(*dto.Update) (*dto.UpdateResponse, error)
}

// update updates bloodrequest
type update struct {
	storeBloodRequest storebloodrequest.BloodRequests
	validate          *validator.Validate
}

func (u *update) toModel(userbloodrequest *dto.Update) (bloodrequest *model.BloodRequest) {
	bloodrequest = &model.BloodRequest{}
	bloodrequest.CreatedAt = time.Now().UTC()
	bloodrequest.UpdatedAt = bloodrequest.CreatedAt
	bloodrequest.UserID = userbloodrequest.UserID
	//	bloodrequest.ID = userbloodrequest.StatusID
	return
}

func (u *update) validateData(update *dto.Update) (err error) {
	err = update.Validate(u.validate)
	return
}

func (u *update) convertData(update *dto.Update) (
	modelBloodRequest *model.BloodRequest,
) {
	modelBloodRequest = u.toModel(update)
	return
}

func (u *update) askStore(modelBloodRequest *model.BloodRequest) (
	id string,
	err error,
) {
	id, err = u.storeBloodRequest.Save(modelBloodRequest)
	return
}

func (u *update) giveResponse(
	modelBloodRequest *model.BloodRequest,
	id string,
) *dto.UpdateResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelBloodRequest.UserID,
	}).Debug("User updated bloodrequest successfully")

	return &dto.UpdateResponse{
		Message:    "BloodRequest updated",
		OK:         true,
		ID:         id,
		UpdateTime: modelBloodRequest.UpdatedAt.String(),
	}
}

func (u *update) giveError() (err error) {
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

//Update implements Update interface
func (u *update) Update(update *dto.Update) (
	*dto.UpdateResponse, error,
) {
	if err := u.validateData(update); err != nil {
		return nil, err
	}

	modelBloodRequest := u.convertData(update)
	id, err := u.askStore(modelBloodRequest)
	if err == nil {
		return u.giveResponse(modelBloodRequest, id), nil
	}

	logrus.Error("Could not update bloodrequest ", err)
	err = u.giveError()
	return nil, err
}

//NewUpdate returns new instance of NewCreate
func NewUpdate(store storebloodrequest.BloodRequests, validate *validator.Validate) Updater {
	return &update{
		store,
		validate,
	}
}
