package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Donor provides dto for donor request
type Donor struct {
	ID           string `json:"donor_request_id"`
	Name         string `json:"name"`
	Phone        string `json:"phone_number"`
	District     string `json:"district"`
	BloodGroup   string `json:"blood_group"`
	Address      string `json:"address"`
	Availability bool   `json:"availability"`
	TimesDonated int    `json:"times_donated"`
	UserID       string `json:"user_id"`
}

//Validate validates donor request data
func (c *Donor) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for donor", false},
			},
		)
	}
	return nil
}

//FromReader reads donor request from request body
func (c *Donor) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(c)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid donor data", false},
		})
	}

	return nil
}
