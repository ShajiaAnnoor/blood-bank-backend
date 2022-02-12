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
type deleteHandler struct {
	delete staticcontent.Creater
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	staticcontent dto.StaticContent,
	err error,
) {
	staticcontent = dto.StaticContent{}
	err = staticcontent.FromReader(body)

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
	staticcontent *dto.StaticContent,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(comment)
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

	staticcontent, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	staticcontent.UserID = ch.decodeContext(r)

	data, err := ch.askController(&staticcontent)

	if err != nil {
		message := "Unable to create staticcontent for status error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for NewCommentRoute
type DeleteParams struct {
	dig.In
	Create     staticcontent.Creater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make comments
func CreateRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.StaticContentCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}