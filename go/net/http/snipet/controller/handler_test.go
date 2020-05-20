package controller

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

var dummyData = map[string]int{
	"budougumi0617": 33,
	"john":          30,
}

type mock struct {
	data map[string]int
}

func (m *mock) GetAge(name string) (int, error) {
	return m.data[name], nil
}

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
		{
			name:  "oka",
			query: "name=john",
			want: HelloResponse{
				Name: "john",
				Age:  30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "http://example.com/hello?"+tt.query, nil)
			w := httptest.NewRecorder()
			m := &mock{}
			m.data = dummyData
			hello := &Hello{}
			hello.repo = m
			hello.HelloHandler(w, req)

			resp := w.Result()

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
