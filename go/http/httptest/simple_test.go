package httptest

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func simple(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Test")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, h)
}

func TestSimple_ByMux(t *testing.T) {
	path := "/simple"
	want := "dummy body"
	req := httptest.NewRequest("GET", path, bytes.NewBufferString(want))
	req.Header.Set("Test", want)

	got := httptest.NewRecorder()

	mu := http.NewServeMux()
	mu.HandleFunc(path, simple)

	mu.ServeHTTP(got, req)

	if got.Body.String() != want {
		t.Errorf("want %s, but got = %s\n", want, got.Body.String())
	}
}
