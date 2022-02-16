package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/organization-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Organization provides dto for organization request
type Organization struct {
	ID           string `json:"organization_request_id"`
	Name         string `json:"name"`
	Phone        string `json:"phone_number"`
	District     string `json:"district"`
	BloodGroup   string `json:"blood_group"`
	Address      string `json:"address"`
	Availability bool   `json:"availability"`
	Status       string `json:"status"`
	TimesDonated int    `json:"times_donated"`
	UserID       string `json:"user_id"`
}

//Validate validates organization request data
func (c *Organization) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for organization", false},
			},
		)
	}
	return nil
}

//FromReader reads organization request from request body
func (c *Organization) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(c)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid organization data", false},
		})
	}

	return nil
}
