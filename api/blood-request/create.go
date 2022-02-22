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

//createHandler holds handler for creating blood requests
type createHandler struct {
	create bloodrequest.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	bloodreqDat dto.BloodReq,
	err error,
) {
	bloodreqDat = dto.BloodReq{}
	err = bloodreqDat.FromReader(body)

	return
}

func (ch *createHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (ch *createHandler) askController(
	bloodreq *dto.BloodReq,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(bloodreq)
	return
}

func (ch *createHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *createHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.CreateResponse,
) {
	routeutils.ServeResponse(
		w,
		http.StatusOK,
		resp,
	)
}

//ServeHTTP implements http.Handler interface. It marks the start of a particular request from the developer perspective
func (ch *createHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	//the follwing line will decode a blood request given by the user
	bloodreqDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	bloodreqDat.UserID = ch.decodeContext(r)

	data, err := ch.askController(&bloodreqDat)

	if err != nil {
		message := "Unable to create blood request error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for CreateRoute
type CreateParams struct {
	dig.In
	Create     bloodreq.Creater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets user create blood requests
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.BloodReqCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
