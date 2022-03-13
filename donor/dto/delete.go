package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// CreateResponse provides create response
type DeleteResponse struct {
	Message     string `json:"message"`
	OK          bool   `json:"ok"`
	ID          string `json:"blood_request_id"`
	RequestTime string `json:"request_time"`
}

// String provides string repsentation
func (dr *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", dr.Message, dr.OK)
}

// Update provides dto for user status update
type Delete struct {
	Comment   string `json:"comment"`
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}

//Validate validates comment update data
func (d *Delete) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid comment update data", false},
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
				Base: errors.Base{"invalid comment update data", false},
			},
		)
	}

	return nil
}
