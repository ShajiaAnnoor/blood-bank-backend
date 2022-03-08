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

//createHandler holds handler for creating notices
type createHandler struct {
	create notice.Creater
}

func (ch *createHandler) decodeBody(
	body io.ReadCloser,
) (
	notice dto.Notice,
	err error,
) {
	notice = dto.Notice{}
	err = notice.FromReader(body)

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
	notice *dto.Notice,
) (
	data *dto.CreateResponse,
	err error,
) {
	data, err = ch.create.Create(notice)
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

	notice, err := ch.decodeBody(r.Body)

	if err != nil {
		message := "Unable to decode error: "
		ch.handleError(w, err, message)
		return
	}

	notice.UserID = ch.decodeContext(r)

	data, err := ch.askController(&notice)

	if err != nil {
		message := "Unable to create notice error: "
		ch.handleError(w, err, message)
		return
	}

	ch.responseSuccess(w, data)
}

//CreateParams provide parameters for NewCommentRoute
type CreateParams struct {
	dig.In
	Create     notice.Creater
	Middleware *middleware.Auth
}

//CreateRoute provides a route that lets users make notices
func CreateRoute(params CreateParams) *routeutils.Route {
	handler := createHandler{params.Create}
	return &routeutils.Route{
		Method:  http.MethodPost,
		Pattern: apipattern.NoticeCreate,
		Handler: params.Middleware.Middleware(&handler),
	}
}
