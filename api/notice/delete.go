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
	delete notice.Deleter
}

func (dh *deleteHandler) decodeBody(
	body io.ReadCloser,
) (
	notice dto.Delete,
	err error,
) {
	err = notice.FromReader(body)
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

func (dh *deleteHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (dh *deleteHandler) askController(update *dto.Delete) (
	resp *dto.DeleteResponse,
	err error,
) {
	resp, err = dh.delete.Delete(update)
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

	noticeDat := dto.Delete{}
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
	Delete     notice.Deleter
	Middleware *middleware.Auth
}

//UpdateRoute provides a route that deletes notice
func DeleteRoute(params DeleteParams) *routeutils.Route {
	handler := deleteHandler{params.Delete}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.NoticeDelete,
		Handler: params.Middleware.Middleware(&handler),
	}
}
