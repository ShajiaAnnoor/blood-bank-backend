package notice

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/notice"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	"go.uber.org/dig"
)

//updateHandler holds notice update handler
type updateHandler struct {
	update notice.Updater
}

func (uh *updateHandler) decodeBody(
	body io.ReadCloser,
) (
	notice dto.Update,
	err error,
) {
	err = notice.FromReader(body)
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

	noticeDat := dto.Update{}
	noticeDat, err := uh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode notice error: "
		uh.handleError(w, err, message)
		return
	}

	noticeDat.UserID = uh.decodeContext(r)

	data, err := uh.askController(&noticeDat)

	if err != nil {
		message := "Unable to update notice error: "
		uh.handleError(w, err, message)
		return
	}

	uh.responseSuccess(w, data)
}

//UpdateParams provide parameters for notice update handler
type UpdateParams struct {
	dig.In
	Update     notice.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates notice
func UpdateRoute(params UpdateParams) *routeutils.Route {
	handler := updateHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.NoticeUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
