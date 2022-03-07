package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Update provides dto for user status update
type Update struct {
	Notice   string `json:"notice"`
	UserID   string `json:"user_id"`
	NoticeID string `json:"notice_id"`
}

//Validate validates notice update data
func (c *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				errors.Base{"invalid notice update data", false},
			},
		)
	}
	return nil
}

//FromReader decodes notice update data from request
func (c *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(c)
	if err != nil {
		return fmt.Errorf(
			"%s:%w",
			err.Error(),
			&errors.Invalid{
				Base: errors.Base{"invalid notice update data", false},
			},
		)
	}

	return nil
}
