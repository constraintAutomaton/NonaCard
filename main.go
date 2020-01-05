package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	middleware "github.com/constraintAutomaton/nonaCard/middleware"
	route "github.com/constraintAutomaton/nonaCard/route"
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
	http.ListenAndServe(getPort(), r)
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
