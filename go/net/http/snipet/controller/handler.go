package controller

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Repository interface {
	GetAge(string) (int, error)
}

type Hello struct {
	repo Repository
}

func NewHello(r Repository) *Hello {
	h := &Hello{}
	h.repo = r
	return h
}
//
//func NewAPIClient(dbName string) *Hello{
//	db sql.Open("sql", dbName)
//	repo = Repo{db}
//	return &Hello{repo}
//}

// クエリパラメータでもらった名前をDBから検索して、年齢を返すハンドラー
func (h *Hello) HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, err := h.repo.GetAge(name)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	resp := HelloResponse{
		Name: name,
		Age:  age,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
