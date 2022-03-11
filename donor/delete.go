package donor

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/donor/dto"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storedonor "gitlab.com/Aubichol/blood-bank-backend/store/donor"
	validator "gopkg.in/go-playground/validator.v9"
)

//Deleter provides an interface for updating noticees
type Deleter interface {
	Delete(*dto.Delete) (*dto.DeleteResponse, error)
}

// delete deletes notice
type delete struct {
	storeDonor storedonor.Donor
	validate   *validator.Validate
}

func (u *delete) toModel(usernotice *dto.Delete) (notice *model.Notice) {
	notice = &model.Notice{}
	notice.CreatedAt = time.Now().UTC()
	//	notice.DeletedAt = notice.CreatedAt
	//	notice.Status = usernotice.Status
	notice.UserID = usernotice.UserID
	notice.ID = usernotice.StatusID
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
	id, err = u.storeDonor.Save(modelNotice)
	return
}

func (u *delete) giveResponse(
	modelNotice *model.Notice,
	id string,
) *dto.DeleteResponse {
	logrus.WithFields(logrus.Fields{
		"id": modelNotice.UserID,
	}).Debug("User deleted notice successfully")

	return &dto.DeleteResponse{
		Message: "Notice deleted",
		OK:      true,
		ID:      id,
		//		DeleteTime: modelNotice.DeletedAt.String(),
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

	logrus.Error("Could not delete notice ", err)
	err = u.giveError()
	return nil, err
}

//NewDelete returns new instance of NewCreate
func NewDelete(store storedonor.Donor, validate *validator.Validate) Deleter {
	return &delete{
		store,
		validate,
	}
}
