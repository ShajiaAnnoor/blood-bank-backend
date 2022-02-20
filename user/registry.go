package user

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storeuser "gitlab.com/Aubichol/blood-bank-backend/store/user"
	"gitlab.com/Aubichol/blood-bank-backend/user/dto"
	"gopkg.in/go-playground/validator.v9"
)

// Registry provides Register method to register a user
type Registry interface {
	Register(register *dto.Register) (*dto.BaseResponse, error)
}

// registry registers user
type registry struct {
	storeUsers storeuser.Users
	validate   *validator.Validate
}

func (r *registry) toModel(register *dto.Register) (user *model.User) {
	user = &model.User{}
	user.FirstName = register.FirstName
	user.LastName = register.LastName
	user.Gender = register.Gender
	user.BirthDate.Year = register.BirthDate.Year
	user.BirthDate.Month = register.BirthDate.Month
	user.BirthDate.Day = register.BirthDate.Day
	user.Email = register.Email
	user.Password = register.Password
	user.Verified = true //TO-DO: it should be made false
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = user.CreatedAt

	return
}

func (r *registry) validateData(register *dto.Register) (err error) {
	err = register.Validate(r.validate)
	return
}

func (r *registry) convertData(register *dto.Register) (
	modelUser *model.User,
) {
	modelUser = r.toModel(register)
	return
}

func (r *registry) askStore(modelUser *model.User) (
	err error,
) {
	err = r.storeUsers.Save(modelUser)
	return
}

func (r *registry) responseError(err error) (
	*dto.BaseResponse,
	error,
) {
	return nil, err
}

func (r *registry) noError(err error) (isError bool) {
	if err == nil {
		isError = true
	}
	return
}

func (r *registry) logSuccess(modelUser *model.User) {
	logrus.WithFields(logrus.Fields{
		"email": modelUser.Email,
		"id":    modelUser.ID,
	}).Debug("User registered")
}

func (r *registry) prepareErrorReponse() (err error) {
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "Invalid data",
		},
	}
	err = fmt.Errorf("%s %w", err.Error(), &errResp)
	return
}

func (r *registry) logError(message string, err error) {
	logrus.Error(message, err)
}

func (r *registry) prepareSuccessfulResponse() (
	resp *dto.BaseResponse,
) {
	resp = &dto.BaseResponse{
		Message: "registered",
		OK:      true,
	}
	return
}

func (r *registry) giveResponse(resp *dto.BaseResponse) (
	*dto.BaseResponse,
	error,
) {
	return resp, nil
}

//Register implements Registry interface
func (r *registry) Register(register *dto.Register) (*dto.BaseResponse, error) {
	if err := r.validateData(register); err != nil {
		return r.responseError(err)
	}

	var modelUser *model.User
	modelUser = r.convertData(register)

	err := r.askStore(modelUser)
	if r.noError(err) {
		r.logSuccess(modelUser)
		resp := r.prepareSuccessfulResponse()
		return r.giveResponse(resp)
	}

	message := "could not save user "
	r.logError(message, err)

	err = r.prepareErrorReponse()
	return r.responseError(err)
}

// NewRegistry returns new instance of NewRegistry
func NewRegistry(storeUsers storeuser.Users, validate *validator.Validate) Registry {
	return &registry{
		storeUsers,
		validate,
	}
}
