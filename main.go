package main

import (
	"fmt"
	"log"
	"net/http"

	middleware "github.com/featTheB/anifan-card/middleware"
	route "github.com/featTheB/anifan-card/route"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	static := r.PathPrefix("/static/")
	fs := http.FileServer(http.Dir("assets/"))

	static.Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", route.Dashboard).Methods("GET")

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprintf(w, "<h1>Welcome the API</h1>")
	}).Methods("GET")

	searchAnime := api.PathPrefix("/search/anime").Subrouter()
	searchAnime.Use(middleware.SearchAnimeAnilist)
	searchAnime.HandleFunc("", route.SearchAnime).Methods("GET")

	getUser := api.PathPrefix("/user").Subrouter()
	getUser.Use(middleware.GetUserInfoAnilist)
	getUser.HandleFunc("/{user}", route.GetUser).Methods("GET")

	log.Println("server started at 8080")
	http.ListenAndServe(":8080", r)
}
