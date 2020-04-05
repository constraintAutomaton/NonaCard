package route

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	private "github.com/constraintAutomaton/nonaCard/rest/routeInterface"
)

// Dashboard generate the main page
func Dashboard(w http.ResponseWriter, r *http.Request) {

	aouthLink := fmt.Sprintf(os.Getenv("OAUTH_LINK"), os.Getenv("CLIENT_ID"), os.Getenv("REDIRECT_URL"))
	cardIdentifier := []string{"card-1", "card-2", "card-3", "card-4", "card-5", "card-6", "card-7", "card-8", "card-9"}
	data := pageSetup{
		CardIds: cardIdentifier,
		Header: private.Header{
			CSS:   "static/css/main.css",
			Title: "NonaCard"},
		Footer: private.Footer{
			Js: "/static/dist/main.js"},
		HeaderDashboard: private.HeaderDashboard{
			OauthLink: aouthLink}}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

var tmpl = template.Must(template.ParseFiles(
	"rest/views/main.html",
	"rest/views/components/header.html",
	"rest/views/components/footer.html",
	"rest/views/components/dashboard/card.html",
	"rest/views/components/dashboard/footer.html",
	"rest/views/components/dashboard/header.html",
))

type pageSetup struct {
	Header          private.Header
	Footer          private.Footer
	HeaderDashboard private.HeaderDashboard
	CardIds         []string
}
