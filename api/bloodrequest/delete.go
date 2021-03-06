package bloodrequest

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"go.uber.org/dig"
)

//deleteHandler holds comment update handler
type deleteHandler struct {
	delete bloodrequest.Deleter
}

func (dh *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	bloodrequest dto.Delete,
	err error,
) {
	err = bloodrequest.FromReader(body)
	return
}

func (dh *deleteHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (dh *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (dh *deleteHandler) askController(update *dto.Delete) (
	resp *dto.DeleteResponse,
	err error,
) {
	resp, err = dh.delete.Delete(update)
	return
}

func (dh *deleteHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.DeleteResponse,
) {
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

//ServeHTTP implements http.Handler interface
func (dh *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	bloodrequest := dto.Delete{}
	bloodrequest, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode bloodrequest delete error: "
		dh.handleError(w, err, message)
		return
	}

	bloodrequest.UserID = dh.decodeContext(r)

	data, err := dh.askController(&bloodrequest)

	if err != nil {
		message := "Unable to delete bloodrequest error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

//DeleteParams provide parameters for bloodrequest delete handler
type DeleteParams struct {
	dig.In
	Delete     bloodrequest.Deleter
	Middleware *middleware.Auth
}

//DeleteRoute provides a route that deletes bloodrequest
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.BloodRequestDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
