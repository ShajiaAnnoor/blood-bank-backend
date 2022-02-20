package user

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
	"gitlab.com/Aubichol/blood-bank-backend/token"
	"gitlab.com/Aubichol/blood-bank-backend/token/dto"
	"go.uber.org/dig"
)

type registerTokenHandler struct {
	registerToken token.Register
}

func (rt *registerTokenHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (rt *registerTokenHandler) askController(
	token dto.Token,
) (
	resp *dto.BaseResponse,
	err error,
) {
	resp, err = rt.registerToken.Register(token)
	return
}

func (rt *registerTokenHandler) responseSuccess(
	w http.ResponseWriter,
	response *dto.BaseResponse,
) {
	routeutils.ServeResponse(w, http.StatusOK, response)
}

func (rt *registerTokenHandler) handleToken(
	w http.ResponseWriter,
	token dto.Token,
) {

	response, err := rt.askController(token)
	if err != nil {
		message := "Unable to register token, err: "
		rt.handleError(w, err, message)
		return
	}

	rt.responseSuccess(w, response)
}

func (rt *registerTokenHandler) decodeBody(
	body io.ReadCloser,
) (
	token dto.Token,
	err error,
) {
	err = token.FromReader(body)
	return
}

//ServeHTTP implements http.Handler interface
func (rt *registerTokenHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close() // needs to add this to other projects as well

	request := dto.Token{}
	var err error
	request, err = rt.decodeBody(r.Body)

	if err != nil {
		logrus.Error("Unable to decode, err: ", err)
		routeutils.ServeError(
			w,
			&errors.Invalid{
				errors.Base{
					"Invalid token", false,
				},
			},
		)
		return
	}
	rt.handleToken(w, request)
}

//OpenRegisterTokenParams lists paramters for
//NewRegisterTokenRoute
type OpenRegisterTokenParams struct {
	dig.In
	RegToken   token.Register
	Middleware *middleware.Auth
}

//NewRegisterTokenRoute provides route for creating a
//registration token
func NewRegisterTokenRoute(
	params OpenRegisterTokenParams,
) *routeutils.Route {
	handler := registerTokenHandler{
		registerToken: params.RegToken,
	}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.RegistrationToken,
		Handler: params.Middleware.Middleware(&handler),
	}
}
