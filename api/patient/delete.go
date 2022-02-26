package patient

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/patient"
	"gitlab.com/Aubichol/blood-bank-backend/patient/dto"
	"go.uber.org/dig"
)

//deleteHandler holds handler for deleting patients
type deleteHandler struct {
	delete patient.Deleter
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	patient dto.Patient,
	err error,
) {
	patient = dto.Patient{}
	err = patient.FromReader(body)

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
	patient *dto.Patient,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = dh.delete.Update(patient)
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
	resp *dto.CreateResponse,
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

	patient, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		dh.handleError(w, err, message)
		return
	}

	patient.UserID = dh.decodeContext(r)

	data, err := dh.askController(&patient)

	if err != nil {
		message := "Unable to delete patient error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//DeleteParams provide parameters for DeleteRoute
type DeleteParams struct 
	dig.In
	Update     patient.Updater
	Middleware *middleware.Auth
}

//DeleteRoute provides a route that lets users make comments
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.PatientUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
