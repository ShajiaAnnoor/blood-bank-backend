package organization

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/api/middleware"
	"gitlab.com/Aubichol/hrishi-backend/api/routeutils"
	"gitlab.com/Aubichol/hrishi-backend/apipattern"
	"gitlab.com/Aubichol/hrishi-backend/comment/dto"
	"go.uber.org/dig"
)

//createHandler holds handler for creating comments
type deleteeHandler struct {
	delete organization.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	organization dto.Organization,
	err error,
) {
	organization = dto.Organization{}
	err = organization.FromReader(body)

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
	organization *dto.Organization,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(organization)
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

	organization, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	organization.UserID = ch.decodeContext(r)

	data, err := ch.askController(&comment)

	if err != nil {
		message := "Unable to create comment for status error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for NewCommentRoute
type DeleteParams struct {
	dig.In
	Create     organization.Creater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make comments
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.OrganizationDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
