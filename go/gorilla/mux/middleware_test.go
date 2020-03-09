package mux

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Ref: https://gist.github.com/maoueh/624f108ee2f3e6ca0b496d6c2f75bcd7
func TestMiddleware(t *testing.T) {
	router := mux.NewRouter()

	router.Use(middleware1)

	wsRouter := router.PathPrefix("/ws").Subrouter()
	wsRouter.Use(middleware2)
	wsRouter.Use(middleware3)

	wsRouter.HandleFunc("/sub", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling ws /sub")
		w.Write([]byte("/sub (ws)"))
	}))

	chainRouter := router.PathPrefix("/chain").Subrouter()
	chainRouter.HandleFunc("/sub1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling chain /sub1")
		w.Write([]byte("/sub1"))
	}))
	chainRouter.HandleFunc("/sub2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling chain /sub2")
		w.Write([]byte("/sub2"))
	}))

	restRouter := router.PathPrefix("/").Subrouter()
	restRouter.Use(middleware3)

	restRouter.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling rest /")
		w.Write([]byte("/ (rest)"))
	}))

	ts := httptest.NewServer(router)
	defer ts.Close()
	tryRequest(t, "GET", ts.URL+"/ws/sub", "/ (sub2)")
	tryRequest(t, "GET", ts.URL+"/", "/ (sub2)")
	tryRequest(t, "GET", ts.URL+"/chain/sub1", "/ (sub2)")
}

func tryRequest(t *testing.T, method, path, want string) {
	cli := &http.Client{}
	r, err := http.NewRequest(method, path, strings.NewReader(""))
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
		t.Errorf("%s: reading body: %v", path, err)
	}

	if string(body) != want {
		t.Errorf("want %q, but got %q", want, body)
	}
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware1")
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware2")
		next.ServeHTTP(w, r)
	})
}

func middleware3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware3")
		next.ServeHTTP(w, r)
	})
}
