package bloodrequest

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"go.uber.org/dig"
)

//updateHandler holds comment update handler
type deleteHandler struct {
	update bloodreq.Updater
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	bloodreq dto.Update,
	err error,
) {
	err = bloodreq.FromReader(body)
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

func (ch *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *deleteHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = ch.update.Update(update)
	return
}

func (ch *deleteHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.UpdateResponse,
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

	bloodreq := dto.Update{}
	bloodreq, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode blood request error: "
		ch.handleError(w, err, message)
		return
	}

	bloodreq.UserID = ch.decodeContext(r)

	data, err := ch.askController(&bloodreq)

	if err != nil {
		message := "Unable to update blood request error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//UpdateParams provide parameters for blood request update handler
type DeleteParams struct {
	dig.In
	Update     bloodreq.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates blood request
func UpdateRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.BloodreqUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
