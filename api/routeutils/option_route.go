package routeutils

import "net/http"

//NewOptionRoute provides a route that adds option routes to all
//the routes
func NewOptionRoute() *Route {

	return &Route{
		Method:  http.MethodOptions,
		Pattern: "/api/v1/*",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ServeResponse(w, http.StatusOK, map[string]interface{}{"message": "ok"})
		}),
	}
}
