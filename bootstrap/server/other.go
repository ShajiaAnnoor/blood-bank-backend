package server

import (
	"gitlab.com/Aubichol/blood-bank-backend/container"
	"gopkg.in/go-playground/validator.v9"
)

//Validator registers validation provider
func Validator(c container.Container) {
	c.Register(validator.New)
}
