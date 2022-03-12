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

//updateHandler holds blood request update handler
type updateHandler struct {
	update bloodreq.Updater
}

func (ch *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	bloodreqAtt dto.Update,
	err error,
) {
	err = bloodreqAtt.FromReader(body)
	return
}

func (ch *updateHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (ch *updateHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *updateHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = ch.update.Update(update)
	return
}

func (ch *updateHandler) responseSuccess(
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
func (ch *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	bloodreqDat := dto.Update{}
	bloodreqDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode blood request update error: "
		ch.handleError(w, err, message)
		return
	}

	bloodreqDat.UserID = ch.decodeContext(r)

	data, err := ch.askController(&bloodreqDat)

	if err != nil {
		message := "Unable to update blood request error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//UpdateParams provide parameters for blood request update handler
type UpdateParams struct {
	dig.In
	Update     bloodreq.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates blood request
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := updateHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.BloodRequestUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
