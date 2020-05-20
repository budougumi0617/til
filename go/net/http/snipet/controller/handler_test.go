package controller

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  HelloResponse
	}{
		{
			name:  "simple",
			query: "name=budougumi0617",
			want: HelloResponse{
				Name: "budougumi0617",
				Age:  33,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://example.com/hello?"+tt.query, nil)
			w := httptest.NewRecorder()
			HelloHandler(w, req)

			resp := w.Result()

			// {"name":"budougumi0617", "age":33}
			var got HelloResponse
			if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
				t.Fatalf("failed decode %q", err)
			}
			if tt.want.Name != got.Name {
				t.Errorf("want %q, but got %q", tt.want.Name, got.Name)
			}
			if tt.want.Age != got.Age {
				t.Errorf("want %d, but got %d", tt.want.Age, got.Age)
			}
		})
	}
}
