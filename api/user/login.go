package user

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/user"
	"gitlab.com/Aubichol/blood-bank-backend/user/dto"
	"go.uber.org/dig"
)

//loginHandler holds login handler
type loginHandler struct {
	userLogin user.Login
}

func (lh *loginHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (lh *loginHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.Token,
) {
	routeutils.ServeResponse(w, http.StatusOK, resp)
}

func (lh *loginHandler) decodeBody(body io.ReadCloser) (
	login dto.Login, err error,
) {
	err = login.FromReader(body)
	return
}

//ServeHTTP implements http.Handler interface for loginHandler
func (lh *loginHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	login := dto.Login{} // Can we move this initialization to dig?
	var err error
	login, err = lh.decodeBody(r.Body)
	if err != nil {
		message := "Unable to decode error: "
		lh.handleError(w, err, message)
		//routeutils.ServeError(w, &errors.Invalid{errors.Base{"Invalid data", false}})
		return
	}

	data, err := lh.userLogin.CreateToken(&login)
	if err != nil {
		message := "Unable to login user error: "
		lh.handleError(w, err, message)
		return
	}
	lh.responseSuccess(w, data)
}

//LoginRouteParams lists all the paramters for NewLoginRoute
type LoginRouteParams struct {
	dig.In
	UserLogin user.Login
}

//LoginRoute provides a route that gives a user login options
func LoginRoute(
	params LoginRouteParams,
) *routeutils.Route {
	handler := loginHandler{params.UserLogin}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.LoginUser,
		Handler: &handler,
	}
}
