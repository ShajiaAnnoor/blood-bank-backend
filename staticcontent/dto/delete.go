package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// DeleteResponse provides create response
type DeleteResponse struct {
	Message    string `json:"message"`
	OK         bool   `json:"ok"`
	ID         string `json:"static_content_id"`
	DeleteTime string `json:"delete_time"`
}

// String provides string repsentation
func (c *DeleteResponse) String() string {
	return fmt.Sprintf("message:%s, ok:%v", c.Message, c.OK)
}

// Update provides dto for user status update
type Delete struct {
	UserID          string `json:"user_id"`
	StaticContentID string `json:"staticcontent_id"`
}

//Validate validates comment update data
func (d *Delete) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid static content update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes organization update data from request
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
