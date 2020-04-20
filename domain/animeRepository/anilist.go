package animerepository

import (
	response "github.com/constraintAutomaton/nonaCard/domain/animeRepository/response"
	remoteCommunication "github.com/constraintAutomaton/nonaCard/domain/remote-communication"
)

const urlAnilist string = "https://graphql.anilist.co"

// Anilist module of the anilist database
type Anilist struct {
	remoteCommunicationModule remoteCommunication.RemoteCommuncation
}

//SearchByName return the result of a search by name in the anilist database
func (anilist Anilist) SearchByName(query string) (response.SearchResponse, error) {
	query := remoteCommunication.ParameterQuery{Url: urlAnilist,
		Query: querySearchAnilist, Variables: map[string]string{
			"name": "query"}, Out: response.SearchAnilistJSON{}, Authorization: make([]string, 0)}

	err := remoteCommunicationModule.Fetch(query)
	if err != nil {
		return response.SearchResponse{}, err
	}
	return anilist.formatSearchResponse(&query.Out), nil
}

func (anilist Anilist) formatSearchResponse(res *response.SearchAnilistJSON) response.SearchResponse {
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
