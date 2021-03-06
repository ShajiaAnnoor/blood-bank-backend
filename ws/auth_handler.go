package ws

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/user"
)

const token = "token"

type authHandler struct {
	sessionVerifier user.SessionVerifier
	clientStore     ClientStore
}

func (ah *authHandler) validate(data *RequestDTO) error {
	if _, ok := data.Values[token]; !ok {
		return errors.New("empty token")
	}

	return nil
}

func (ah *authHandler) Handle(c Client, data *RequestDTO) {
	if err := ah.validate(data); err != nil {
		logrus.Error("validate: ", err)
		c.Kick()
		return
	}

	session, err := ah.sessionVerifier.VerifySession(data.Values[token])
	if err != nil {
		logrus.Error("verify session: ", err)
		c.Kick()
		return
	}

	if session == nil {
		c.Kick()
		return
	}

	c.SetID(session.UserID)
	ah.clientStore.Add(c)
}

func NewAuthHandler(
	sessionVerifier user.SessionVerifier,
	clientStore ClientStore,
) Handler {
	return &authHandler{
		sessionVerifier: sessionVerifier,
		clientStore:     clientStore,
	}
}
