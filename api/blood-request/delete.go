package bloodrequest

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	bloodreq "gitlab.com/Aubichol/blood-bank-backend/bloodrequest"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"go.uber.org/dig"
)

//deleteHandler holds blood request delete handler
type deleteHandler struct {
	delete bloodreq.Updater
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
	resp, err = ch.delete.Update(update)
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

	bloodreqDat := dto.Update{}
	bloodreqDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode blood request error: "
		ch.handleError(w, err, message)
		return
	}

	bloodreq.UserID = ch.decodeContext(r)

	data, err := ch.askController(&bloodreqDat)

	if err != nil {
		message := "Unable to update blood request error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//DeleteParams provide parameters for blood request delete handler
type DeleteParams struct {
	dig.In
	Delete     bloodreq.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates blood request
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.BloodreqUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
