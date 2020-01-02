package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/constraintAutomaton/nonaCard/pkg/clientGraphQl"
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
	}
  }
  `

// GetUserInfoAnilist get the user info from the anilist api
func GetUserInfoAnilist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		variables := map[string]string{
			"user": vars["user"]}
		var res UserAnilistJSON

		m, err := clientGraphQl.Fetch(urlAnilist, queryGetUserInfoAnilist, variables, &res)
		if err != nil {
			log.Println(err)
			log.Println(m)

		} else {
			b, err := json.Marshal(res.formatAnilistUserInfo())
			if err == nil {
				w.Header().Set("Content-Type", "text/json; application/json")
				io.WriteString(w, string(b))
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (u UserAnilistJSON) formatAnilistUserInfo() FormatedAnilistUserInfo {
	user := u.Data.User
	statistics := user.Statistics.Anime
	tags := statistics.Tags
	formatedTag := make([]TagsFormated, len(tags), len(tags))
	for i, tag := range tags {
		formatedTag[i] = TagsFormated{Count: tag.Count,
			Name:        tag.Tag.Description,
			Description: tag.Tag.Description}
	}
	res := FormatedAnilistUserInfo{
		SiteURL: user.SiteURL,
		Avatar:  user.Avatar,
		Statistics: statisticsFormated{MeanScore: statistics.MeanScore,
			StandardDeviation: statistics.StandardDeviation,
			MinutesWatched:    statistics.MinutesWatched,
			Count:             statistics.Count,
			Scores:            statistics.Scores,
			Tags:              formatedTag}}
	return res
}
