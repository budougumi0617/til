package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type internal struct {
	called bool
}

func (i *internal) RoundTrip(req *http.Request) (*http.Response, error) {
	i.called = true
	base := http.DefaultTransport
	return base.RoundTrip(req)
}

type transport struct {
	Base http.RoundTripper
}

func (t *transport) base() http.RoundTripper {
	if t.Base == nil {
		return http.DefaultTransport
	}
	return t.Base
}

// RoundTrip is custom RoundTrip.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// do anything...
	return t.base().RoundTrip(req)
}

func TestRoundTripper(t *testing.T) {
	rt := &internal{}
	cli := &http.Client{
		Transport: &transport{
			Base: rt,
		},
	}
	p := "/example"
	mux := http.NewServeMux()
	b := "want body"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, b)
	})
	pl := "payload"
	TryRequest(t, "check RoundTripper", "GET", p, pl, mux, http.StatusOK, b, cli)
	if !rt.called {
		t.Fatal("interior RoundTripper is not called")
	}

}

// https://medium.com/@timakin/go-api-testing-173b97fb23ec
func TryRequest(t *testing.T, desc, method, path, payload string, mux *http.ServeMux, wantCode int, wantBody string, c *http.Client) {
	srv := httptest.NewServer(mux)
	defer srv.Close()

	req, err := http.NewRequest(method, srv.URL+path, strings.NewReader(payload))
	if err != nil {
		t.Errorf("%s: generate request: %v", desc, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

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
