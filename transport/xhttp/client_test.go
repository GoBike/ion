package xhttp

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

/*
Basically, its pretty repetitive when it comes to testing a http-client.
  1. Setup a mocked-server.
  2. Create a bunch of handlers. this will be your expected output.
  3. RPC to mocked-server, via mocked-server-url. - input
*/

// plainTextBodyHandler writes plain-text data w/o fancy encoding. i.e.: json.
type plainTextBodyHandler string

func (b plainTextBodyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(b))
}

// start a mocked server, add a handler that writes plain-text into
// http-response's Body. Expect decoder returns same plain-text as response.
func TestBody(t *testing.T) {

	var (
		testbody = "testbody"
		enc      = func(context.Context, *http.Request, interface{}) error { return nil }
		dec      = func(_ context.Context, r *http.Response) (interface{}, error) {

			buffer := make([]byte, len(testbody))
			_, err := r.Body.Read(buffer)

			if err != nil && err != io.EOF {
				return nil, err
			}
			return string(buffer), nil
		}
	)

	ts := httptest.NewServer(plainTextBodyHandler(testbody))
	defer ts.Close()
	u, _ := url.Parse(ts.URL)

	httpclient := NewClient("", u, enc, dec)

	resp, err := httpclient.Rpc()(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp != testbody {
		t.Errorf("want: %s, got: %s", testbody, resp)
	}
}
