package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// DeleteResponse provides delete response
type DeleteResponse struct {
	Message     string `json:"message"`
	OK          bool   `json:"ok"`
	ID          string `json:"blood_request_id"`
	RequestTime string `json:"request_time"`
}

// String provides string repsentation
func (d *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", d.Message, d.OK)
}

// Delete provides dto for blood request delete
type Delete struct {
	UserID         string `json:"user_id"`
	BloodRequestID string `json:"blood_request_id"`
	IsDeleted      bool   `json:"is_deleted"`
}

//Validate validates blood request delete data
func (d *Delete) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid blood request delete data", false},
			},
		)
	}
	return nil
}

//FromReader decodes comment update data from request
func (d *Delete) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(d)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid blood request delete data", false},
			},
		)
	}

	return nil
}
