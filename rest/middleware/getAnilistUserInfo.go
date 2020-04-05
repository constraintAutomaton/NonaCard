package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	graphClient "github.com/constraintAutomaton/nonaCard/rest/graphQl/client"
	"github.com/gorilla/mux"
)

// GetUserInfoAnilist get the user info from the anilist api
func GetUserInfoAnilist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		variables := map[string]string{
			"user": vars["user"]}
		var res UserAnilistJSON

		res.getUserInfoAPI(&w, &variables)
		formatResponseUserAnilist(&w, &res)
		next.ServeHTTP(w, r)
	})
}
func (res *UserAnilistJSON) getUserInfoAPI(w *http.ResponseWriter, variables *map[string]string) {
	err := graphClient.Fetch(urlAnilist, queryGetUserInfoAnilist, variables, &res)
	if err != nil {
		log.Println(err)
		http.Error(*w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Fprintf(*w, "")
	}
}
func formatResponseUserAnilist(w *http.ResponseWriter, res *UserAnilistJSON) {
	b, err := json.Marshal(res.getFormatAnilistUserInfo())
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
func (res UserAnilistJSON) getFormatAnilistUserInfo() FormatedAnilistUserInfo {
	user := res.Data.User
	statistics := user.Statistics.Anime
	tags := statistics.Tags
	formatedTag := make([]TagsFormated, len(tags))
	for i, tag := range tags {
		formatedTag[i] = TagsFormated{Count: tag.Count,
			Name:        tag.Tag.Description,
			Description: tag.Tag.Description}
	}
	return FormatedAnilistUserInfo{
		SiteURL: user.SiteURL,
		Avatar:  user.Avatar,
		Statistics: statisticsFormated{MeanScore: statistics.MeanScore,
			StandardDeviation: statistics.StandardDeviation,
			MinutesWatched:    statistics.MinutesWatched,
			Count:             statistics.Count,
			Scores:            statistics.Scores,
			Tags:              formatedTag}}
}

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
