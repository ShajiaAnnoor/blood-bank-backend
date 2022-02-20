package user

import (
	"fmt"

	"gitlab.com/Aubichol/blood-bank-backend/errors"
	storeuser "gitlab.com/Aubichol/blood-bank-backend/store/user"
	"gitlab.com/Aubichol/blood-bank-backend/user/dto"
)

//MyProfiler is an interface for delivering profile info
type MyProfiler interface {
	Me(string) (*dto.Me, error)
}

//MyProfileFunc is a type that implements MyProfiler
type MyProfileFunc func(string) (*dto.Me, error)

//Me implements MyProfiler interface
func (m MyProfileFunc) Me(id string) (*dto.Me, error) {
	return m(id)
}

//NewMyProfile provides a MyProfiler interface
func NewMyProfile(storeUsers storeuser.Users) MyProfiler {
	f := func(id string) (*dto.Me, error) {
		user, err := storeUsers.FindByID(id)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", err.Error(), &errors.Unknown{
				errors.Base{"Could not find user", false},
			})
		}

		if user == nil {
			return nil, &errors.Unknown{
				errors.Base{"could not find user", false},
			}
		}
		me := dto.Me{}
		me.FromModel(user)
		if len(me.ProfilePic) > 0 {
			me.ProfilePic = "/api/v1/pictures/" + me.ProfilePic
		}
		return &me, nil
	}

	return MyProfileFunc(f)
}
