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

//deleteHandler holds handler for creating patients
type deleteHandler struct {
	create patient.Creater
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

func (ch *deleteHandler) askController(
	patient *dto.Patient,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(patient)
	return
}

func (ch *deleteHandler) decodeContext(
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
func (ch *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	patient, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	patient.UserID = ch.decodeContext(r)

	data, err := ch.askController(&patient)

	if err != nil {
		message := "Unable to delete patient error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for NewCommentRoute
type DeleteParams struct 
	dig.In
	Update     patient.Updater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make comments
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.PatientUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
