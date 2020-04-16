package response

type SearchResponse struct {
	Result []result `json:"data"`
}

// result information about the media query
type result struct {
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
