package main

import (
	"log"
	"net/http"

	"github.com/budougumi0617/til/go/net/http/snipet/repository"
	"github.com/budougumi0617/til/go/net/http/snipet/controller"
)

func main() {
	mux := http.NewServeMux()
	// DBへの接続を開く
	repo := repository.NewRepository()
	hello := controller.NewHello(repo)

	mux.HandleFunc("/hello", hello.HelloHandler)

	s := http.Server{
		Addr:    ":18080",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	}
}
