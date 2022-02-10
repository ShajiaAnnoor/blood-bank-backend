package staticcontent

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

//createHandler holds handler for creating comments
type createHandler struct {
	create staticcontent.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	staticcontent dto.StaticContent,
	err error,
) {
	staticcontent = dto.StaticContent{}
	err = staticcontent.FromReader(body)

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
	staticcontent *dto.StaticContent,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(comment)
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

//ServeHTTP implements http.Handler interface
func (ch *createHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	comment, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	staticcontent.UserID = ch.decodeContext(r)

	data, err := ch.askController(&comment)

	if err != nil {
		message := "Unable to create staticcontent for status error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for NewCommentRoute
type CreateParams struct {
	dig.In
	Create     staticcontent.Creater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make comments
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.StaticContentCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
