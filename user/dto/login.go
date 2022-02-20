package dto

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
)

//Login stores login related data
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//FromReader decodes to json type data from request
func (l *Login) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(l)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"Invalid login data", false},
		})
	}

	// Need to refactor password validation to a separate function

	if l.Password == "" {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"Empty password", false},
		})
	}

	h := md5.New()
	_, err = io.WriteString(h, l.Password)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Unknown{
			Base: errors.Base{"Could not convert password to hash", false},
		})
	}
	l.Password = fmt.Sprintf("%x", h.Sum(nil))

	//TO-DO: need to verify email as well
	return nil
}

//Token stores token data
type Token struct {
	Token string `json:"token"`
}
