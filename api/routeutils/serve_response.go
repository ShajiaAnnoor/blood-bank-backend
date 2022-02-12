package routeutils

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/blood-bank-backend/errors"
)

//ServeResponse performs routine tasks for sending a response
func ServeResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		logrus.Error("could not serve http response ", err)
	}
}

//ServeError responds when an error takes place
func ServeError(w http.ResponseWriter, err error) {
	data, ok := errors.HTTPUnwrap(err)
	if !ok {
		ServeResponse(
			w,
			http.StatusInternalServerError,
			map[string]interface{}{
				"message": "internal server error",
				"ok":      false,
			},
		)
		return
	}

	ServeResponse(w, data.Status, data.Data)
}
