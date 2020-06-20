package middleware

import (
	"bytes"
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rb := &bytes.Buffer{}
		// TODO: wをラップする
		next.ServeHTTP(w, r)
		log.Printf("resp:= %q", rb)
	})
}