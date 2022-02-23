package donor

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/donor"
	"gitlab.com/Aubichol/blood-bank-backend/donor/dto"
	"go.uber.org/dig"
)

//deleteHandler holds comment update handler
type deleteHandler struct {
	delete donor.Updater
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

func (dh *deleteHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = dh.delete.Update(update)
	return
}

func (dh *deleteHandler) responseSuccess(
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
func (dh *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	donor := dto.Update{}
	donor, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode donor delete error: "
		dh.handleError(w, err, message)
		return
	}

	donor.UserID = dh.decodeContext(r)

	data, err := dh.askController(&donor)

	if err != nil {
		message := "Unable to delete donor error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

//DeleteParams provide parameters for donor delete handler
type DeleteParams struct {
	dig.In
	Delete     donor.Updater
	Middleware *middleware.Auth
}

//DeleteRoute provides a route that deletes donor
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.DonorUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
