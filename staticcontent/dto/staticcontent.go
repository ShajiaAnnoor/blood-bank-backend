package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// StaticContent provides dto for staticcontent request
type StaticContent struct {
	ID     string `json:"staticcontent_request_id"`
	Text   string `json:"name"`
	UserID string `json:"user_id"`
}

//Validate validates staticcontent request data
func (sc *StaticContent) Validate(validate *validator.Validate) error {
	if err := validate.Struct(sc); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for static content", false},
			},
		)
	}
	return nil
}

//FromReader reads staticcontent request from request body
func (sc *StaticContent) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(sc)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid staticcontent data", false},
		})
	}

	return nil
}
