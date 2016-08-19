package xhttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Basically, its pretty repetitive when it comes to testing a http-client.
  1. Setup a mocked-server, along with handlers. - expected output
  2. RPC to mocked-server, with mocked-server-url. - input
*/

func TestIntegration(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello client!")
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)

	if err != nil {
		log.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", greeting)

}
