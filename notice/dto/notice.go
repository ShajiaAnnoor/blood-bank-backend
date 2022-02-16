package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Notice provides dto for notice request
type Notice struct {
	ID          string `json:"notice_request_id"`
	PatientName string `json:"patient_name"`
	BloodGroup  string `json:"notice_group"`
	Description string `json:"description"`
	District    string `json:"district"`
	Address     string `json:"address"`
	Title       string `json:"title"`
	UserID      string `json:"user_id"`
}

//Validate validates notice request data
func (c *Notice) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for notice", false},
			},
		)
	}
	return nil
}

//FromReader reads notice request from request body
func (c *Notice) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(c)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid notice data", false},
		})
	}

	return nil
}
