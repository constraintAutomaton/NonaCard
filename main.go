package main

import (
	"fmt"
	"net/http"

	dashboard "github.com/featTheB/anifan-card/route"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	static := r.PathPrefix("/static/")
	fs := http.FileServer(http.Dir("assets/"))

	static.Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", dashboard.Page).Methods("GET")

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome the API")
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
