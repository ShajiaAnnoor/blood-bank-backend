package organization

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	"go.uber.org/dig"
)

//createHandler holds handler for creating comments
type deleteHandler struct {
	delete organization.Deleter
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	organization dto.Organization,
	err error,
) {
	organization = dto.Organization{}
	err = organization.FromReader(body)

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
	organization *dto.Organization,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.delete.Update(organization)
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

	organization, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	organization.UserID = ch.decodeContext(r)

	data, err := ch.askController(organization)

	if err != nil {
		message := "Unable to create organization error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for NewCommentRoute
type DeleteParams struct {
	dig.In
	Update     organization.Updater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make comments
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.OrganizationDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
