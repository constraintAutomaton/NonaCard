package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	graphClient "github.com/constraintAutomaton/nonaCard/rest/graphQl/client"
)

// SearchAnimeAnilist search an anime in anilist
func SearchAnimeAnilist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		variables := map[string]string{
			"name": q}
		var res SearchAnilistJSON
		res.getSearchQuery(&w, &variables)
		formatResponseSearchAnilist(&w, &res)
		next.ServeHTTP(w, r)

	})
}
func (res *SearchAnilistJSON) getSearchQuery(w *http.ResponseWriter, variables *map[string]string) {
	err := graphClient.Fetch(urlAnilist, querySearchAnilist, variables, &res)
	if err != nil {
		log.Println(err)
		http.Error(*w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Fprintf(*w, "")
	}
}
func formatResponseSearchAnilist(w *http.ResponseWriter, res *SearchAnilistJSON) {
	b, err := json.Marshal((*res).Data.Page.Media)
	if err != nil {
		log.Println(err)
		http.Error(*w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Fprintf(*w, "")
	}
	_, err = io.WriteString(*w, string(b))
	if err != nil {
		log.Println(err)
		http.Error(*w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Fprintf(*w, "")
	}
	(*w).Header().Set("Content-Type", "text/json; application/json")
}

const querySearchAnilist = `query ($name: String) {
	Page {
	  media(search: $name, type: ANIME) {
		id
		idMal
		averageScore
		siteUrl
		title {
		  romaji
		  english
		}
		coverImage {
		  medium
		  large
		  color
		}
		description
		tags{
			description
			name
		  }
	  }
	}
  }`
