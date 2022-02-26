package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/hrishi-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Update provides dto for user status update
type Update struct {
	Patient   string `json:"patient"`
	UserID    string `json:"user_id"`
	PatientID string `json:"patient_id"`
}

//Validate validates patient update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid patient update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes patient update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid patient update data", false},
			},
		)
	}

	return nil
}
