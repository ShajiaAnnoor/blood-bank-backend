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
	ID           string `json:"donor_id"`
	Name         string `json:"name"`
	Phone        string `json:"phone_number"`
	District     string `json:"district"`
	BloodGroup   string `json:"blood_group"`
	Address      string `json:"address"`
	Availability bool   `json:"availability"`
	TimesDonated int    `json:"times_donated"`
	UserID       string `json:"user_id"`
}

//Validate validates donor update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid donor update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes donor update data from request
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid donor update data", false},
			},
		)
	}

	return nil
}
