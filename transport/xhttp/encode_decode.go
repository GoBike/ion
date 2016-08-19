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
