package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type rwWrapper struct {
	rw  http.ResponseWriter
	out io.Writer
}

func NewRwWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{
		rw:  rw,
		out: io.MultiWriter(rw, buf),
	}
}

func (r *rwWrapper) Header() http.Header {
	return r.rw.Header()
}

func (r *rwWrapper) Write(i []byte) (int, error) {
	return r.out.Write(i)
}

func (r *rwWrapper) WriteHeader(statusCode int) {
	r.rw.WriteHeader(statusCode)
}

// 実際はloggerをDIするほうだろう。
func NewLogger(out io.Writer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rww := NewRwWrapper(w, out)
			next.ServeHTTP(rww, r)
		})
	}
}

func Test(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestLogger(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello, client")
	})
	buf := &bytes.Buffer{}
	l := NewLogger(buf)
	ts := httptest.NewServer(l(h))
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
	if !bytes.Equal(want, buf.Bytes()) {
		t.Errorf("want %q, but %q", want, buf)
	}
}
