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

func (dh *deleteHandler) askController(update *dto.Update) (
	resp *dto.UpdateResponse,
	err error,
) {
	resp, err = dh.update.Update(update)
	return
}

func (dh *deleteHandler) responseSuccess(
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
func (dh *deleteHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	noticeDat := dto.Update{}
	noticeDat, err := dh.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode notice error: "
		dh.handleError(w, err, message)
		return
	}

	noticeDat.UserID = dh.decodeContext(r)

	data, err := dh.askController(&noticeDat)

	if err != nil {
		message := "Unable to update notice error: "
		dh.handleError(w, err, message)
		return
	}

	dh.responseSuccess(w, data)
}

//DeleteParams provide parameters for notice delete handler
type DeleteParams struct {
	dig.In
	Delete     notice.Updater
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that deletes notice
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.NoticeUpdate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
