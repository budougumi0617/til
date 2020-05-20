package main

import (
	"log"
	"net/http"

	"github.com/budougumi0617/til/go/net/http/snipet/controller"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", controller.HelloHandler)

	s := http.Server{
		Addr:    ":18080",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	}
}
