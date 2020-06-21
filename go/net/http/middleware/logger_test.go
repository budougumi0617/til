package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type rwWrapper struct {
	rw http.ResponseWriter
	mw io.Writer
}

func NewRwWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{
		rw: rw,
		mw: io.MultiWriter(rw, buf),
	}
}

func (r *rwWrapper) Header() http.Header {
	return r.rw.Header()
}

func (r *rwWrapper) Write(i []byte) (int, error) {
	return r.mw.Write(i)
}

func (r *rwWrapper) WriteHeader(statusCode int) {
	r.rw.WriteHeader(statusCode)
}

func NewLogger(l *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf := &bytes.Buffer{}
			rww := NewRwWrapper(w, buf)
			next.ServeHTTP(rww, r)
			l.Printf("%s", buf)
		})
	}
}

func TestLogger(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello, client")
	})
	buf := &bytes.Buffer{}
	l := log.New(buf, "", 0)
	ts := httptest.NewServer(NewLogger(l)(h))
	t.Cleanup(ts.Close)

	cli := &http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequestWithContext(context.TODO(), "GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// compare logging data and client received value.
	want, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
	// \nが混ざっているので、完全一致にはならない
	if !strings.Contains(buf.String(), string(want)) {
		t.Errorf("want %q, but %q", want, buf)
	}
}
