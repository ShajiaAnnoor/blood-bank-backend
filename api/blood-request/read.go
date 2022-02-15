package bloodrequest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/api/middleware"
	"gitlab.com/Aubichol/blood-bank-backend/api/routeutils"
	"gitlab.com/Aubichol/blood-bank-backend/apipattern"
	"gitlab.com/Aubichol/blood-bank-backend/bloodrequest/dto"
	"go.uber.org/dig"
)

type readHandler struct {
	reader bloodreq.Reader
}

func (read *readHandler) decodeURL(
	r *http.Request,
) (donorID string) {
	// Get user id from url
	donorID = chi.URLParam(r, "id")
	return
}

func (read *readHandler) decodeContext(
	r *http.Request,
) (userID string) {
	userID = r.Context().Value("userID").(string)
	return
}

func (read *readHandler) askController(
	req *dto.ReadReq,
) (
	resp *dto.ReadResp,
	err error,
) {
	resp, err = read.reader.Read(req)
	return
}

func (read *readHandler) handleError(
	w http.ResponseWriter,
	err error,
) {
	logrus.Error(err)
	routeutils.ServeError(w, err)
}

func (read *readHandler) responseSuccess(
	w http.ResponseWriter,
	resp *dto.ReadResp,
) {
	// Serve a response to the client
	routeutils.ServeResponse(w, http.StatusOK, resp)
}

func (read *readHandler) handleRead(
	w http.ResponseWriter,
	r *http.Request,
) {

	req := dto.ReadReq{}
	req.BloodreqID = read.decodeURL(r)

	req.BloodreqID = read.decodeContext(r)

	// Read request from database using request id and user id
	resp, err := read.askController(&req)

	if err != nil {
		read.handleError(w, err)
		return
	}

	read.responseSuccess(w, resp)
}

//ServeHTTP implements http.Handler
func (read *readHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	defer r.Body.Close()

	read.handleRead(w, r)
}

//ReadRouteParams lists all the parameters for ReadRoute
type ReadRouteParams struct {
	dig.In
	Reader     bloodreq.Reader
	Middleware *middleware.Auth
}

//ReadRoute provides a route to get comment
func ReadRoute(params ReadRouteParams) *routeutils.Route {

	handler := readHandler{
		reader: params.Reader,
	}

	return &routeutils.Route{
		Method:  http.MethodGet,
		Pattern: apipattern.BloodreqRead,
		Handler: params.Middleware.Middleware(&handler),
	}
}
