package token

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storetoken "gitlab.com/Aubichol/blood-bank-backend/store/token"
	"gitlab.com/Aubichol/blood-bank-backend/token/dto"
	validator "gopkg.in/go-playground/validator.v9"
)

// Register provides register method for registering firebase token
type Register interface {
	Register(register dto.Token) (*dto.BaseResponse, error)
}

// register registers firebase token
type register struct {
	storeToken storetoken.Token
	validate   *validator.Validate
}

func (r *register) toModel(registerToken dto.Token) (token *model.Token) {
	token = &model.Token{}
	token.CreatedAt = time.Now().UTC()
	token.UpdatedAt = token.CreatedAt
	token.Token = registerToken.Token
	token.UserID = registerToken.UserID
	return
}

//Register implements Registry interface
func (r *register) Register(register dto.Token) (*dto.BaseResponse, error) {
	if err := register.Validate(r.validate); err != nil {
		return nil, err
	}

	modelToken := r.toModel(register)

	err := r.storeToken.Save(modelToken)
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"id": modelToken.UserID,
		}).Debug("token registered successfully")

		return &dto.BaseResponse{
			Message: "token registered",
			OK:      true,
		}, nil
	}

	logrus.Error("could not register token ", err)
	errResp := errors.Unknown{
		Base: errors.Base{
			OK:      false,
			Message: "invalid data",
		},
	}

	err = fmt.Errorf("%s %w", err.Error(), &errResp)
	return nil, err
}

//NewRegisterStore returns new instance of NewRegister
func NewRegisterStore(storeTokens storetoken.Token, validate *validator.Validate) Register {
	return &register{
		storeTokens,
		validate,
	}
}
