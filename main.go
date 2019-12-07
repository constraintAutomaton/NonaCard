package main

import (
	"net/http"

	dashboard "github.com/featTheB/anilist-user-analysis/route"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test", dashboard.Page)

	http.ListenAndServe(":8080", r)
}
