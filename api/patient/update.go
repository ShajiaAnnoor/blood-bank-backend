package patient

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

//updateHandler holds patient update handler
type updateHandler struct {
	update patient.Updater
}

func (uh *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	patient dto.Update,
	err error,
) {
	err = patient.FromReader(body)
	return
}

func (uh *updateHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (uh *updateHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (uh *updateHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = uh.update.Update(update)
	return
}

func (uh *updateHandler) responseSuccess(
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
func (uh *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	patient := dto.Update{}
	patient, err := uh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode patient error: "
		uh.handleError(w, err, message)
		return
	}

	patient.UserID = uh.decodeContext(r)

	data, err := uh.askController(&patient)

	if err != nil {
		message := "Unable to update patient error: "
		uh.handleError(w, err, message)
		return
	}

	uh.responseSuccess(w, data)
}

//UpdateParams provide parameters for patient update handler
type UpdateParams struct {
	dig.In
	Update     patient.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates patient
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := updateHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.PatientUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
