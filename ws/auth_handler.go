package ws

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/user"
)

const token = "token"

type authHandler struct {
	sessionVerifier user.SessionVerifier
	clientStore     ClientStore
}

func (a *authHandler) validate(data *RequestDTO) error {
	if _, ok := data.Values[token]; !ok {
		return errors.New("empty token")
	}

	return nil
}

func (a *authHandler) Handle(c Client, data *RequestDTO) {
	if err := a.validate(data); err != nil {
		logrus.Error("validate: ", err)
		c.Kick()
		return
	}

	session, err := a.sessionVerifier.VerifySession(data.Values[token])
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
	a.clientStore.Add(c)
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
