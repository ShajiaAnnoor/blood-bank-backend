package donor

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/comment/dto"
	"go.uber.org/dig"
)

//updateHandler holds comment update handler
type deleteHandler struct {
	update donor.Updater
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	donor dto.Update,
	err error,
) {
	err = donor.FromReader(body)
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

	donor := dto.Update{}
	donor, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode donor delete error: "
		ch.handleError(w, err, message)
		return
	}

	donor.UserID = ch.decodeContext(r)

	data, err := ch.askController(&donor)

	if err != nil {
		message := "Unable to delete donor error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//UpdateParams provide parameters for comment update handler
type UpdateParams struct {
	dig.In
	Update     donor.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates comment
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := deleteHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.DonorUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
