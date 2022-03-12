package staticcontent

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent"
	"gitlab.com/Aubichol/blood-bank-backend/staticcontent/dto"
	"go.uber.org/dig"
)

//deleteHandler holds handler for deleting static contents
type deleteHandler struct {
	delete staticcontent.Deleter
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	staticcontent dto.Delete,
	err error,
) {
	staticcontent = dto.Delete{}
	err = staticcontent.FromReader(body)

	return
}

func (ch *deleteHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (dh *deleteHandler) askController(
	staticcontent *dto.Delete,
) (
	data *dto.DeleteResponse,
	err error,
) {
	data, err = dh.delete.Delete(staticcontent)
	return
}

func (dh *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *deleteHandler) responseSuccess(
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
func (ch *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	staticcontent, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	staticcontent.UserID = ch.decodeContext(r)

	data, err := ch.askController(&staticcontent)

	if err != nil {
		message := "Unable to update staticcontent error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//DeleteParams provide parameters for DeleteRoute
type DeleteParams struct {
	dig.In
	Delete     staticcontent.Deleter
	Middleware *middleware.Auth
}

//DeleteRoute provides a route that lets users make static contents
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.StaticContentUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
