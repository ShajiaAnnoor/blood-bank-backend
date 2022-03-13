package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Update provides dto for notice update
type Update struct {
	Notice   string `json:"notice"`
	UserID   string `json:"user_id"`
	NoticeID string `json:"notice_id"`
}

//Validate validates notice update data
func (u *Update) Validate(validate *validator.Validate) error {
	if err := validate.Struct(u); err != nil {
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
func (u *Update) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(u)
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
