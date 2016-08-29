package xhttp

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

// Most test req/resp scheduled to run at 1 millisecond mark.
// Test expects to fail after TimeOut
const TimeOut = 1 * time.Millisecond

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
func TestDecodePlainTextResponse(t *testing.T) {

	var (
		testbody = "testbody"
		enc      = NopEncodeRequest()
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
	turl, _ := url.Parse(ts.URL)

	httpclient := NewClient("", turl, enc, dec)

	resp, err := httpclient.Rpc()(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp != testbody {
		t.Errorf("want: %s, got: %s", testbody, resp)
	}
}

// start a mocked server, client send a request without header, expects no error.
func TestNoBefore(t *testing.T) {
	fmt.Println("todo: TestNoBefore")
}

// start a mocked server, client send a request with 2 before option. expects
// server to receive them both.
func TestTwoBefore(t *testing.T) {
	fmt.Println("todo: TestTwoBefore")
}

// start mocked server, SetRequestHeader to before, client send request.
// exptects server receives header info.
func TestOneBefore(t *testing.T) {

	var (
		eavesdrop = make(chan string, 1)
		key       = "x-gobike-foo"
		value     = "bar"
		enc       = NopEncodeRequest()
		dec       = NopDecodeResponse()
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eavesdrop <- r.Header.Get(key)
	}))
	defer ts.Close()

	turl, _ := url.Parse(ts.URL)

	client := NewClient("", turl, enc, dec, SetBefore(SetRequestHeader(key, value)))
	_, err := client.Rpc()(context.Background(), nil)

	if err != nil {
		t.FailNow()
	}

	select {
	case got := <-eavesdrop:
		if want := value; got != want {
			t.FailNow()
		}
	case <-time.After(TimeOut):
		t.FailNow()
	}
}
