package xhttp

import (
	"context"
	"net/http"
)

// RequestFunc is a generic callback to intercept a HTTP request.
//
// In Servers, RequestFuncs are executed prior to invoking the endpoint.
// In Clients, RequestFuncs are executed after creating request but prior
// to invoking the HTTP client (Do request).
type RequestFunc func(context.Context, *http.Request) context.Context

// ClientResponseFunc may take information from an HTTP request and make the
// response available for consumption. ClientResponseFuncs are only executed in
// clients, after a request has been made, but prior to it being decoded.
type ClientResponseFunc func(context.Context, *http.Response) context.Context

// SetRequestHeader supercharges RequestFunc with additional input: HTTP header
// key-value pair.
func SetRequestHeader(key string, value string) RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		r.Header.Set(key, value)
		return ctx
	}
}
