package user

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/model"
	storeuser "gitlab.com/Aubichol/blood-bank-backend/store/user"
	"gitlab.com/Aubichol/blood-bank-backend/user/dto"
)

//EmailAndPasswordChecker is an interface that checks the validity for email and password
type EmailAndPasswordChecker interface {
	EmailAndPasswordCheck(*dto.Login) (*model.User, error)
}

//EmailAndPasswordCheckerFunc implements EmailAndPasswordChecker
type EmailAndPasswordCheckerFunc func(*dto.Login) (*model.User, error)

//EmailAndPasswordCheck implements EmailAndPasswordChecker interface
func (e EmailAndPasswordCheckerFunc) EmailAndPasswordCheck(dto *dto.Login) (*model.User, error) {
	return e(dto)
}

//NewEmailAndPasswordChecker provides EmailAndPasswordChecker
func NewEmailAndPasswordChecker(storeUsers storeuser.Users) EmailAndPasswordChecker {
	f := func(login *dto.Login) (*model.User, error) {
		user, err := storeUsers.FindByEmail(login.Email)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err":   err,
				"email": login.Email,
			}).Error("to get user by email")

			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid email or password", false},
			}
		}

		if user == nil {
			logrus.WithField("email", login.Email).Debug("Got empty user")
			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid email or password", false},
			}
		}

		if user.Password != login.Password {
			logrus.Debug("Passwords did not match")
			return nil, &errors.Unknown{
				Base: errors.Base{"Invalid email or password", false},
			}
		}
		return user, nil
	}

	return EmailAndPasswordCheckerFunc(f)
}
