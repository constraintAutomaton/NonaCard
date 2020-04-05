package middleware

// UserAnilistJSON struct of the json for the user of anilist
type UserAnilistJSON struct {
	Data user `json:"data"`
}
type user struct {
	User userInfo `json:"User"`
}
type userInfo struct {
	SiteURL    string     `json:"siteUrl"`
	Avatar     avatar     `json:"avatar"`
	Statistics statistics `json:"statistics"`
}
type avatar struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
}
type statistics struct {
	Anime animeInfo `json:"anime"`
}
type animeInfo struct {
	MeanScore         float64    `json:"meanScore"`
	StandardDeviation float64    `json:"standardDeviation"`
	MinutesWatched    float64    `json:"minutesWatched"`
	Count             int        `json:"count"`
	Scores            []scores   `json:"scores"`
	Tags              []tagsUser `json:"tags"`
}
type scores struct {
	Score float64 `json:"score"`
	Count int     `json:"count"`
}
type tagsUser struct {
	Tag   Tags
	Count int `json:"count"`
}

// FormatedAnilistUserInfo user data sended to the user
type FormatedAnilistUserInfo struct {
	SiteURL    string             `json:"siteUrl"`
	Avatar     avatar             `json:"avatar"`
	Statistics statisticsFormated `json:"statistics"`
}
type statisticsFormated struct {
	MeanScore         float64        `json:"meanScore"`
	StandardDeviation float64        `json:"standardDeviation"`
	MinutesWatched    float64        `json:"minutesWatched"`
	Count             int            `json:"count"`
	Scores            []scores       `json:"scores"`
	Tags              []TagsFormated `json:"tags"`
}

// TagsFormated tag with a simpler format for alleviate calculation in the frontend
type TagsFormated struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}
