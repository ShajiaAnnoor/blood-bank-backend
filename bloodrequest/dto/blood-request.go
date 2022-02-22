package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// BloodReq provides dto for blood request
type BloodReq struct {
	ID         string `json:"blood_request_id"`
	Request    string `json:"request"`
	BloodGroup string `json:"blood_group"`
	UserID     string `json:"user_id"`
}

//Validate validates blood request data
func (b *BloodReq) Validate(validate *validator.Validate) error {
	if err := validate.Struct(b); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid data for blood request", false},
			},
		)
	}
	return nil
}

//FromReader reads blood request from request body
func (b *BloodReq) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(b)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid blood request data", false},
		})
	}

	return nil
}
