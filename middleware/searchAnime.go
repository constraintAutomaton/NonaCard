package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/featTheB/anifan-card/pkg/clientGraphQl"
)

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

// SearchAnimeAnilist search an anime in anilist
func SearchAnimeAnilist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		variables := map[string]string{
			"name": q}
		var res SearchAnilistJSON

		m, err := clientGraphQl.Fetch(urlAnilist, querySearchAnilist, variables, &res)
		if err != nil {
			log.Println(err)
			log.Println(m)

		} else {
			b, err := json.Marshal(res.Data.Page.Media)
			if err == nil {
				w.Header().Set("Content-Type", "text/json; application/json")
				io.WriteString(w, string(b))
			}
		}

		next.ServeHTTP(w, r)

	})
}
