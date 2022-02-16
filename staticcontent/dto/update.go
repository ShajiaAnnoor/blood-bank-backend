package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/hrishi-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Update provides dto for user status update
type Update struct {
	Comment   string `json:"comment"`
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}

//Validate validates comment update data
func (c *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
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
func (c *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(c)
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
