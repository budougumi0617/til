package mux

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Ref: https://gist.github.com/maoueh/624f108ee2f3e6ca0b496d6c2f75bcd7
func TestMiddleware(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("body"))
	}
	middleware1 := func(buf *bytes.Buffer) func(next http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf.Write([]byte("1"))
				next.ServeHTTP(w, r)
			})
		}
	}

	middleware2 := func(buf *bytes.Buffer) func(next http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf.Write([]byte("2"))
				next.ServeHTTP(w, r)
			})
		}
	}

	middleware3 := func(buf *bytes.Buffer) func(next http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf.Write([]byte("3"))
				next.ServeHTTP(w, r)
			})
		}
	}

	tests := [...]struct {
		name       string
		router     func(*bytes.Buffer) *mux.Router
		path, want string
	}{
		{
			name: "simple",
			path: "/sub/sub",
			router: func(buf *bytes.Buffer) *mux.Router {
				r := mux.NewRouter()
				sub := r.PathPrefix("/sub").Subrouter()
				sub.HandleFunc("/sub", h)
				// あとからUseしても適応される。
				sub.Use(middleware2(buf), middleware3(buf))
				// rootのルーターは全体に適応される。
				// ただし、後からでもsub routerより優先される。
				r.Use(middleware1(buf))
				return r
			},
			want: "123",
		},
		{
			name: "AccessRoot",
			path: "/",
			router: func(buf *bytes.Buffer) *mux.Router {
				r := mux.NewRouter()
				r.HandleFunc("/", h)
				r.Use(middleware1(buf))
				sub := r.PathPrefix("/sub").Subrouter()
				// 子routerのMiddlewareは親routerに適応されない。
				sub.Use(middleware2(buf))
				return r
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			ts := httptest.NewServer(tt.router(buf))
			defer ts.Close()
			cli := &http.Client{}
			r, err := http.NewRequest("GET", ts.URL+tt.path, strings.NewReader(""))
			if err != nil {
				t.Errorf("NewRequest failed: %v", err)
			}

			resp, err := cli.Do(r)
			if err != nil {
				t.Errorf("Do failed: %v", err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("%s: reading body: %v", tt.path, err)
			}

			if string(body) != "body" {
				t.Errorf("want 'body', but got %q", body)
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("want %q, but got %q", tt.want, got)
			}
		})
	}
}
