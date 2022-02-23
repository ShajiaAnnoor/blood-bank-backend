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

//updateHandler holds organization update handler
type updateHandler struct {
	update organization.Updater
}

func (ch *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	organization dto.Update,
	err error,
) {
	err = organization.FromReader(body)
	return
}

func (ch *updateHandler) handleError(
	w http.ResponseWriter,
	err error,
	message string,
) {
	logrus.Error(message, err)
	routeutils.ServeError(w, err)
}

func (ch *updateHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *updateHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = ch.update.Update(update)
	return
}

func (ch *updateHandler) responseSuccess(
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
func (ch *updateHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	organization := dto.Update{}
	organization, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode organization error: "
		ch.handleError(w, err, message)
		return
	}

	organization.UserID = ch.decodeContext(r)

	data, err := ch.askController(&organization)

	if err != nil {
		message := "Unable to update organization for user error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//UpdateParams provide parameters for organization update handler
type UpdateParams struct {
	dig.In
	Update     organization.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates comment
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := updateHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.OrganizationUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
