package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	middleware "github.com/constraintAutomaton/nonaCard/middleware"
	route "github.com/constraintAutomaton/nonaCard/route"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
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

	http.ListenAndServe(getPort(), r)
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	} else {
		log.Println("Starting app at ", port)
	}
	return ":" + port
}
