package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// Update provides dto for user status update
type Update struct {
	StaticContent   string `json:"staticcontent"`
	UserID          string `json:"user_id"`
	StaticContentID string `json:"staticcontent_id"`
}

//Validate validates staticcontent update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid staticcontent update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes staticcontent update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid staticcontent update data", false},
			},
		)
	}

	return nil
}
