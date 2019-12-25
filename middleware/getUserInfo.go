package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/featTheB/anifan-card/pkg/clientGraphQl"
	"github.com/gorilla/mux"
)

const queryGetUserInfoAnilist = `query UserInfo($user: String) {
	User(name: $user) {
	  siteUrl
	  avatar {
		large
		medium
	  }
	  statistics {
		anime {
		  meanScore
		  standardDeviation
		  minutesWatched
		  count
		  scores {
			score
			count
		  }
		  tags {
			tag {
			  name
			  description
			}
			count
		  }
		}
	  }
	  favourites {
		anime {
		  edges {
			node {
			  averageScore
			  title {
				userPreferred
			  }
			  siteUrl
			  coverImage {
				medium
				large
			  }
			  description
			}
		  }
		}
	  }
	}
  }
  `

// GetUserInfoAnilist get the user info from the anilist api
func GetUserInfoAnilist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		variables := map[string]string{
			"user": vars["user"]}
		res := new(map[string]interface{})
		m, err := clientGraphQl.Fetch(urlAnilist, queryGetUserInfoAnilist, variables, res)
		if err != nil {
			log.Println(err)
			log.Println(m)

		} else {
			b, err := json.Marshal((*res)["data"].(map[string]interface{})["User"])
			if err == nil {
				w.Header().Set("Content-Type", "text/json; application/json")
				io.WriteString(w, string(b))
			}
		}
		next.ServeHTTP(w, r)
	})
}
