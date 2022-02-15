package notice

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/notice/dto"
	"go.uber.org/dig"
)

//deleteHandler holds notice update handler
type deleteHandler struct {
	update notice.Updater
}

func (ch *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	notice dto.Update,
	err error,
) {
	err = notice.FromReader(body)
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

func (ch *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (ch *deleteHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = ch.update.Update(update)
	return
}

func (ch *deleteHandler) responseSuccess(
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
func (ch *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	noticeDat := dto.Update{}
	noticeDat, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode notice error: "
		ch.handleError(w, err, message)
		return
	}

	noticeDat.UserID = ch.decodeContext(r)

	data, err := ch.askController(&noticeDat)

	if err != nil {
		message := "Unable to update notice error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//DeleteParams provide parameters for notice update handler
type DeleteParams struct {
	dig.In
	Update     notice.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that updates comment
func UpdateRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Update}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.NoticeUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
