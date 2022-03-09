package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// Patient provides dto for patient request
type Patient struct {
	ID         string `json:"patient_request_id"`
	Name       string `json:"name"`
	BloodGroup string `json:"blood_group"`
	District   string `json:"district"`
	Phone      string `json:"phone_number"`
	Address    string `json:"address"`
	UserID     string `json:"user_id"`
}

//Validate validates patient request data
func (p *Patient) Validate(validate *validator.Validate) error {
	if err := validate.Struct(p); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for patient", false},
			},
		)
	}
	return nil
}

//FromReader reads patient request from request body
func (p *Patient) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(p)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid patient data", false},
		})
	}

	return nil
}
