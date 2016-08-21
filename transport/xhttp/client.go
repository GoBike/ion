package xhttp

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"

	"github.com/gobike/ion/api"
)

// A Client resembles a single remote-procedure-call.
//
// It carries endpoint information, constructs api.Api that is capable of firing
// http-request via a HTTP client.
type Client struct {

	// client is a HTTP client.
	client *http.Client

	// endpoint is a parsed URL. It knows where to hit a f***ing server.
	endpoint *url.URL

	// method specifies HTTP verb.
	method string

	enc    EncodeRequestFunc
	dec    DecodeResponseFunc
	before []RequestFunc
}

// NewClient contructs an usable Client for a single remote method.
func NewClient(method string, endpoint *url.URL, enc EncodeRequestFunc, dec DecodeResponseFunc) *Client {

	c := &Client{
		endpoint: endpoint,
		method:   method,
		enc:      enc,
		dec:      dec,
		before:   []RequestFunc{},
	}

	return c
}

// SetBefore sets >= 1 RequestFunc f to before
func (c *Client) SetBefore(f ...RequestFunc) {
	c.before = f
}

// SetClient sets HTTP client.
func (c *Client) SetClient(client *http.Client) {
	c.client = client
}

// Rpc supercharges/upgrades api.Api with a http-gun, which turns it into a
// badass-RPC.
//
// Now, it can fire http-request.
func (c Client) Rpc() api.Api {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var req *http.Request
		var resp *http.Response

		var err error

		// first, make sure its cancellable.
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		// second, aim to your target
		if req, err = http.NewRequest(c.method, c.endpoint.String(), nil); err != nil {
			return nil, Error{PhaseDo, err}
		}

		// third, fillup bullets
		if err = c.enc(ctx, req, request); err != nil {
			return nil, Error{PhaseEncode, err}
		}
		for _, f := range c.before {
			ctx = f(ctx, req)
		}

		// next, fire!
		if resp, err = ctxhttp.Do(ctx, c.client, req); err != nil {
			return nil, Error{PhaseDo, err}
		}

		// todo: with buffered-stream, this is false.
		defer resp.Body.Close()

		// finally, obtain output.
		var response interface{}
		if response, err = c.dec(ctx, resp); err != nil {
			return nil, Error{PhaseDecode, err}
		}

		return response, nil
	}
}
