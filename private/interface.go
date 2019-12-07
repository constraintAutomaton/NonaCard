package private

const URL string = "https://graphql.anilist.co"

type Media struct {
	id           int
	title        Title
	type_        string
	popularity   int
	episodes     int
	genres       []string
	volumes      int
	source       string
	trailer      Trailer
	description  string
	seasonYear   int
	duration     int
	bannerImage  string
	averageScore int
	trending     int
	studios      Studio
	stats        Stats
}

type Title struct {
	english string
}

type Trailer struct {
	id        int
	thumbnail string
	site      string
}
type Studio struct {
	edge []Edge
}
type Edge struct {
	node Node
}
type Node struct {
	name string
}
type Stats struct {
	scoreDistribution []ScoreDistribution
}
type ScoreDistribution struct {
	score  int
	amount int
}
