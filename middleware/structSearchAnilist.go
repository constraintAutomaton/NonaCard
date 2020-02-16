package middleware

// SearchAnilistJSON struct of the json for the search query of anilist
type SearchAnilistJSON struct {
	Data pageSearchAnilist `json:"data"`
}
type pageSearchAnilist struct {
	Page mediaSearchAnilistAnime `json:"Page"`
}
type mediaSearchAnilistAnime struct {
	Media []MediaAnilist `json:"media"`
}

// MediaAnilist information about the media query
type MediaAnilist struct {
	ID           int        `json:"id"`
	IDMal        int        `json:"idMal"`
	AverageScore float64    `json:"averageScore"`
	SiteURL      string     `json:"siteUrl"`
	Title        title      `json:"title"`
	CoverImage   coverImage `json:"coverImage"`
	Description  string     `json:"description"`
	Tags         []Tags     `json:"tags"`
}
type title struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}
type coverImage struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
	Color  string `json:"color"`
}

//Tags of the anime
type Tags struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
