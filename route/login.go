package route

import (
	"log"
	"net/http"
)

// Login log or register the user
func Login(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "", http.StatusSeeOther)
	token := r.URL.Query()["code"][0]
	connect(token)
}

func connect(pToken string) {
	log.Println(pToken)
}
