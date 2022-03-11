package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
)

//Me stores personal profile related information
type Me struct {
	ID         string                 `json:"id"`
	FirstName  string                 `json:"first_name"`
	LastName   string                 `json:"last_name"`
	Gender     string                 `json:"gender"`
	BirthDate  BirthDate              `json:"birth_date"`
	Email      string                 `json:"email"`
	Profile    map[string]interface{} `json:"profile"`
	ProfilePic string                 `json:"profile_pic"`
}

//FromModel converts model data to json format data
func (m *Me) FromModel(user *model.User) {
	m.ID = user.ID
	m.FirstName = user.FirstName
	m.LastName = user.LastName
	m.Gender = user.Gender
	m.BirthDate.FromModel(&user.BirthDate)
	m.Email = user.Email
	m.Profile = user.Profile
	m.ProfilePic = user.ProfilePic
}

//MeUpdate stores personal profile update related data
type MeUpdate struct {
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Gender    string                 `json:"gender"`
	BirthDate BirthDate              `json:"birth_date"`
	Profile   map[string]interface{} `json:"profile"`
}

//ToModel converts json data to model data
func (m *MeUpdate) ToModel(user *model.User) {
	user.FirstName = m.FirstName
	user.LastName = m.LastName
	user.Gender = m.Gender
	m.BirthDate.ToModel(&user.BirthDate)
	user.Profile = m.Profile
}

//FromReader decodes request data to json type data
func (m *MeUpdate) FromReader(reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(m)
	if err != nil {
		return fmt.Errorf("%s:%w", err.Error(), &errors.Invalid{
			Base: errors.Base{"invalid update data", false},
		})
	}

	return nil
}
