package organization

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	storeorganization "gitlab.com/Aubichol/blood-bank-backend/store/organization"
	"gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for updating organizationes
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes organization
type delete struct {
	storeNotice storeorganization.Notice
	validate    *validator.Validate
}

func (u *delete) toModel(userorganization *dto.Delete) (organization *model.Notice) {
	organization = &model.Notice{}
	organization.CreatedAt = time.Now().UTC()
	organization.DeletedAt = organization.CreatedAt
	organization.Status = userorganization.Status
	organization.UserID = userorganization.UserID
	organization.ID = userorganization.StatusID
	return
}

func (u *delete) validateData(delete *dto.Delete) (err error) {
	err = delete.Validate(u.validate)
	return
}

func (u *delete) convertData(delete *dto.Delete) (
	modelNotice *model.Notice,
) {
	modelNotice = u.toModel(delete)
	return
}

func (u *delete) askStore(modelNotice *model.Notice) (
	id string,
	err error,
) {
	id, err = u.storeNotice.Save(modelNotice)
	return
}

func (u *delete) giveResponse(
	modelNotice *model.Notice,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelNotice.UserID,
	}).Debug("User deleted organization successfully")

	return &dto.DeleteResponse{
		Message:    "Notice deleted",
		OK:         true,
		ID:         id,
		DeleteTime: modelNotice.DeletedAt.String(),
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

	modelNotice := u.convertData(delete)
	id, err := u.askStore(modelNotice)
	if err == nil {
		return u.giveResponse(modelNotice, id), nil
	}

	logrus.Error("Could not delete organization ", err)
	err = u.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storeorganization.Notice, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
