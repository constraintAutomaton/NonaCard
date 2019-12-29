package route

import (
	"html/template"
	"net/http"

	"github.com/featTheB/anifan-card/private"
)

var tmpl = template.Must(template.ParseFiles(
	"views/main.html",
	"views/components/header.html",
	"views/components/footer.html",
	"views/components/card.html"))

type pageSetup struct {
	Header  private.Header
	Footer  private.Footer
	Title   string
	CardIds []string
}

// Dashboard generate the main page
func Dashboard(w http.ResponseWriter, r *http.Request) {
	cardIdentifier := []string{"card-1", "card-2", "card-3", "card-4", "card-5", "card-6", "card-7", "card-8", "card-9"}
	data := pageSetup{
		Title:   "3 By 3 of Anon",
		CardIds: cardIdentifier,
		Header: private.Header{
			CSS:   "static/css/main.css",
			Title: "3By3"},
		Footer: private.Footer{
			Js: "/static/dist/main.js"}}
	tmpl.Execute(w, data)
}
