package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	middleware "github.com/constraintAutomaton/nonaCard/middleware"
	route "github.com/constraintAutomaton/nonaCard/route"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	initializeFrontEnd()
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
	r.HandleFunc("/login", route.Login).Methods("GET")

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

	err := http.ListenAndServe(getPort(), r)
	if err != nil {
		log.Fatal(err)
	}

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

func initializeFrontEnd() {
	cmd := exec.Command("git", "submodule update --init --recursive")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("make", "run-frontEnd")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
