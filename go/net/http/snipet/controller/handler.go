package controller

import (
	"io"
	"net/http"
)

type HelloResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// クエリパラメータでもらった名前をDBから検索して、年齢を返すハンドラー
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"name":"budougumi0617", "age": 50}`)
}
