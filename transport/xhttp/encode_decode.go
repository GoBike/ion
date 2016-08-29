package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
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

// JSONEncodeRequest is a ion/xhttp.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func JSONEncodeRequest() EncodeRequestFunc {
	return func(_ context.Context, r *http.Request, request interface{}) error {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(request); err != nil {
			return err
		}
		r.Body = ioutil.NopCloser(&buf)
		return nil
	}
}
