package organization

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/organization"
	"gitlab.com/Aubichol/blood-bank-backend/organization/dto"
	"go.uber.org/dig"
)

//createHandler holds handler for creating comments
type deleteHandler struct {
	delete organization.Deleter
}

func (dh *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	organization dto.Delete,
	err error,
) {
	organization = dto.Delete{}
	err = organization.FromReader(body)

	return
}

func (dh *deleteHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (dh *deleteHandler) askController(
	organization *dto.Delete,
) (
	data *dto.DeleteResponse,
	err error,
) {
	data, err = dh.delete.Delete(organization)
	return
}

func (dh *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (dh *deleteHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.DeleteResponse,
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

	organization, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		dh.handleError(w, err, message)
		return
	}

	organization.UserID = dh.decodeContext(r)

	data, err := dh.askController(&organization)

	if err != nil {
		message := "Unable to create organization error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

//CreateParams provide parameters for DeleteRoute
type DeleteParams struct {
	dig.In
	Delete     organization.Deleter
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make organizations
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.OrganizationDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
