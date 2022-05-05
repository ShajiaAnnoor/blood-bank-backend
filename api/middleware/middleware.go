package middleware

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/user"
)

//Auth holds the middleware that verifies the session
type Auth struct {
	userSessionVerifier user.SessionVerifier
}

//Middleware implments a middleware that checkes the user session for authorization
func (a *Auth) Middleware(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		session, err := a.userSessionVerifier.VerifySession(token)
		if err != nil {
			logrus.Error("Unable to verify user token, error :", err)
			routeutils.ServeError(w, err)
			return
		}

		if session == nil {
			routeutils.ServeResponse(
				w,
				http.StatusForbidden,
				map[string]interface{}{
					"message": "Invalid token",
					"ok":      false,
				},
			)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(f)
}

//NewAuthMiddleware returns an AuthMiddleware
func NewAuthMiddleware(userSessionVerifier user.SessionVerifier) *Auth {
	return &Auth{userSessionVerifier}
}

//AuthMiddlewareURL stores an user session verifier
type AuthMiddlewareURL struct {
	userSessionVerifier user.SessionVerifier
}

//Middleware implements user token verification
func (a *AuthMiddlewareURL) Middleware(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		tokenParams := r.URL.Query()["token"]
		if len(tokenParams) == 0 {
			routeutils.ServeResponse(w, http.StatusForbidden, map[string]interface{}{
				"message": "Invalid token",
				"ok":      false,
			})
			return
		}

		token := tokenParams[0]
		session, err := a.userSessionVerifier.VerifySession(token)
		if err != nil {
			logrus.Error("Unable to verify user token, error :", err)
			routeutils.ServeError(w, err)
			return
		}

		if session == nil {
			routeutils.ServeResponse(w, http.StatusForbidden, map[string]interface{}{
				"message": "Invalid token",
				"ok":      false,
			})
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(f)
}

//NewAuthMiddlewareURL provides a middlewareurl
func NewAuthMiddlewareURL(userSessionVerifier user.SessionVerifier) *AuthMiddlewareURL {
	return &AuthMiddlewareURL{userSessionVerifier}
}
