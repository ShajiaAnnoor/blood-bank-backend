package user

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/user"
	"gitlab.com/Aubichol/blood-bank-backend/user/dto"
)

//registrationHandler holds registration handler
type registrationHandler struct {
	registry user.Registry
}

func (rh *registrationHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.BaseResponse,
) {
	routeutils.ServeResponse(w, http.StatusOK, resp)
}

func (rh *registrationHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

//ServeHTTP implements http.Handler interface
func (rh *registrationHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	register := dto.Register{}
	if err := register.FromReader(r.Body); err != nil {
		message := "Unable to decode err: "
		rh.handleError(w, err, message)
		return
	}

	data, err := rh.registry.Register(&register)
	if err != nil {
		message := "Unable to register user err: "
		rh.handleError(w, err, message)
		return
	}

	rh.responseSuccess(w, data)
}

//RegistrationRoute provides a route that registers users
func RegistrationRoute(
	userRegistry user.Registry,
) *routeutils.Route {
	handler := registrationHandler{userRegistry}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.UserRegistration,
		Handler: &handler,
	}
}
