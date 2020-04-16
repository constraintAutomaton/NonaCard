package animeRepository

import (
	response "github.com/constraintAutomaton/nonaCard/domain/animeRepository/response"
	graphClient "github.com/constraintAutomaton/nonaCard/rest/graphQl/client"
)

const urlAnilist string = "https://graphql.anilist.co"

type Anilist struct{}

func (this Anilist) SearchByName(query string) (response.SearchResponse, error) {
	variables := map[string]string{
		"name": "query"}

	res := response.SearchAnilistJSON{}
	err := graphClient.Fetch(urlAnilist, querySearchAnilist, &variables, &res)
	if err != nil {
		return response.SearchResponse{}, err
	}
	return this.formatSearchResponse(&res), nil
}

func (this Anilist) formatSearchResponse(res *response.SearchAnilistJSON) response.SearchResponse {
	resFormated := response.SearchResponse{Result: res.Data.Page.Media}
	return resFormated
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
