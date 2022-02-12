package routeutils

import "net/http"

//Route defines a basic route structure
type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}
