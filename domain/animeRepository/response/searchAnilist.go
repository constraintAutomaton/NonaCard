package response

// SearchAnilistJSON struct of the json for the search query of anilist
type SearchAnilistJSON struct {
	Data pageSearchAnilist `json:"data"`
}
type pageSearchAnilist struct {
	Page mediaSearchAnilistAnime `json:"Page"`
}
type mediaSearchAnilistAnime struct {
	Media []result `json:"media"`
}
