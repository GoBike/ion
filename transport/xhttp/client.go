package xhttp

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gobike/ion/api"
)

// A Client resembles a single remote-procedure-call.
//
// It carries endpoint information, constructs api.Api that is capable of firing
// http-request via a HTTP client.
type Client struct {

	// client is a HTTP client.
	client *http.Client

	// endpoint is a parsed URL. It knows where to hit a server.
	endpoint *url.URL

	// method specifies HTTP verb.
	verb string
}

// NewClient contructs an usable Client for a single remote method.
func NewClient(verb string, endpoint *url.URL) *Client {

	c := &Client{
		endpoint: endpoint,
		verb:     verb,
	}

	return c
}

// Rpc supercharges/upgrades api.Api into a badass-RPC, capable of firing
// http-request.
func (c Client) Rpc() api.Api {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}
