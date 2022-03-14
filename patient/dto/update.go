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
	Patient    string `json:"patient"`
	ID         string `json:"patient_id"`
	Name       string `json:"name"`
	BloodGroup string `json:"blood_group"`
	District   string `json:"district"`
	Phone      string `json:"phone_number"`
	Address    string `json:"address"`
	UserID     string `json:"user_id"`
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
