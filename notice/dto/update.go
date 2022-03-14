package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Update provides dto for notice update
type Update struct {
	Notice      string `json:"notice"`
	ID          string `json:"notice_id"`
	PatientName string `json:"patient_name"`
	BloodGroup  string `json:"notice_group"`
	Description string `json:"description"`
	District    string `json:"district"`
	Address     string `json:"address"`
	Title       string `json:"title"`
	UserID      string `json:"user_id"`
}

//Validate validates notice update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid notice update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes notice update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid notice update data", false},
			},
		)
	}

	return nil
}
