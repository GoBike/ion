package xhttp

import (
	"context"
	"net/http"
)

// EncodeRequestFunc is a callback that allows developer to `massage` inputs
// into http.Request, alongside with additional info from context.Context,
// if necessarily.
type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error

// DecodeResponseFunc is a callback that allows developer to `massage` http.Response
// into outputs, alongside with additional info from context.Context,
// if necessarily.
type DecodeResponseFunc func(context.Context, *http.Response) (interface{}, error)

// NopEncodeRequest superchages EncodeRequestFunc with NOP!
func NopEncodeRequest() EncodeRequestFunc {
	return func(context.Context, *http.Request, interface{}) error {
		return nil
	}
}

// NopDecodeResponse superchages DecodeResponseFunc with NOP!
func NopDecodeResponse() DecodeResponseFunc {
	return func(context.Context, *http.Response) (interface{}, error) {
		return nil, nil
	}
}

// StatusOKDecodeResponse is a ion./xhttp.DecodeRequestFunc that
// check if statuscode == 200
func StatusOKDecodeResponse() DecodeResponseFunc {
	return func(_ context.Context, r *http.Response) (interface{}, error) {
		return r.StatusCode == http.StatusOK, nil
	}
}
