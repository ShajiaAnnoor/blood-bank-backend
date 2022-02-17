package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Update provides dto for blood request update
type Update struct {
	Request    string `json:"request"`
	BloodGroup string `json:"blood_group"`
	UserID     string `json:"user_id"`
}

//Validate validates blood request update data
func (c *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid blood request update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes comment update data from request
func (c *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(c)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid blood request update data", false},
			},
		)
	}

	return nil
}
