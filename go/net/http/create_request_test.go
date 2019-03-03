package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	cookie1 := &http.Cookie{
		Name:  "cn1",
		Value: "cv1",
	}
	cookie2 := &http.Cookie{
		Name:  "cn2",
		Value: "cv2",
	}
	type args struct {
		method  string
		status  int
		body    string
		code    string
		cookies []*http.Cookie
	}
	tests := []struct {
		name        string
		args        args
		wantStatus  int
		wantBody    string
		wantCookies []*http.Cookie
	}{
		{
			name: "StatusBadRequest",
			args: args{
				method: "GET",
				status: http.StatusBadRequest,
				body:   "ng",
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "ng",
		},
		{
			name: "StatusOK",
			args: args{
				method: "GET",
				status: http.StatusOK,
				code:   "hogehoge", // temp value
				body:   "ok body",
				cookies: []*http.Cookie{
					cookie1,
					cookie2,
				},
			},
			wantStatus: http.StatusOK,
			wantBody:   "ok body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.args.method, "http://dummy.url.com/", bytes.NewBufferString(tt.args.body))

			// Build queries.
			q := req.URL.Query()
			q.Add("code", tt.args.code)
			req.URL.RawQuery = q.Encode()

			req = req.WithContext(context.Background())

			if len(tt.args.cookies) != 0 {
				for _, c := range tt.args.cookies {
					req.AddCookie(c)
				}
			}

			got := httptest.NewRecorder()

			simple := func(w http.ResponseWriter, r *http.Request) {
				var buf bytes.Buffer
				buf.ReadFrom(r.Body)
				w.WriteHeader(tt.args.status)
				io.WriteString(w, buf.String())
			}

			h := http.HandlerFunc(simple)
			h.ServeHTTP(got, req)
			if got.Code != tt.wantStatus {
				t.Errorf("want %d, but %d", tt.wantStatus, got.Code)
			}
			if got := got.Body.String(); got != tt.wantBody {
				t.Errorf("want %s, but %s", tt.wantBody, got)
			}
			if gots := got.Result().Cookies(); tt.wantCookies != nil {
				for i, got := range gots {
					want := tt.wantCookies[i]
					if got.Name != want.Name || got.Value != want.Value {
						t.Errorf("want %v, but got %v", want, got)
					}
				}
			}
		})
	}
}
