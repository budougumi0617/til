package http

import (
	"io/ioutil"
	"net/http"
	ghttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type transport struct {
	Base ghttp.RoundTripper
}

func (t *transport) base() ghttp.RoundTripper {
	if t.Base == nil {
		return ghttp.DefaultTransport
	}
	return t.Base
}

// RoundTrip is custom RoundTrip.
func (t *transport) RoundTrip(req *ghttp.Request) (*ghttp.Response, error) {
	// do anything...
	return t.base().RoundTrip(req)
}

func TestRoundTripper(t *testing.T) {
	// TODO write test

}

// https://medium.com/@timakin/go-api-testing-173b97fb23ec
func TryRequest(t *testing.T, desc, method, path, payload string, mux *ghttp.ServeMux, wantCode int, wantBody string) {
	srv := httptest.NewServer(mux)
	defer srv.Close()

	req, err := http.NewRequest(method, srv.URL+path, strings.NewReader(payload))
	if err != nil {
		t.Errorf("%s: generate request: %v", desc, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	c := http.DefaultClient

	resp, err := c.Do(req)
	if err != nil {
		t.Errorf("%s: http.Get: %v", desc, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("%s: reading body: %v", desc, err)
		return
	}

	if resp.StatusCode != wantCode {
		t.Errorf("%s: got HTTP %d, want %d", desc, resp.StatusCode, wantCode)
		t.Errorf("response body: %s", string(body))
		return
	}

	if wantBody != "" && string(body) != wantBody {
		t.Errorf("%s: got HTTP body %q, want %q", desc, body, wantBody)
		return
	}
}
