package errors

import (
	"errors"
	"net/http"
)

type HTTPUnwrapValue struct {
	Data   interface{}
	Status int
}

func HTTPUnwrap(err error) (*HTTPUnwrapValue, bool) {
	var invalid *Invalid
	if errors.As(err, &invalid) {
		return &HTTPUnwrapValue{invalid, http.StatusBadRequest}, true
	}

	var unknown *Unknown
	if errors.As(err, &unknown) {
		return &HTTPUnwrapValue{unknown, http.StatusForbidden}, true
	}

	var unauthorized *Unauthorized
	if errors.As(err, &unauthorized) {
		return &HTTPUnwrapValue{unauthorized, http.StatusUnauthorized}, true
	}

	return nil, false
}
