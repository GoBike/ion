package api

import "context"

// Api is a fundamental building block of servers and clients.
// It represents a single RPC method. It is a callback function signature.
type Api func(ctx context.Context, request interface{}) (response interface{}, err error)
