package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	validator "gopkg.in/go-playground/validator.v9"
)

// Token provides dto for firebase token
type Token struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

//Validate validates registration token data
func (t *Token) Validate(validate *validator.Validate) error {
	if err := validate.Struct(t); err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			errors.Base{"Invalid data", false},
		})
	}
	return nil
}

//FromReader converts data from request
func (t *Token) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(t)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"Invalid token data", false},
		})
	}

	return nil
}
