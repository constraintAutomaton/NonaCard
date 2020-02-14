package private

// Header parameter of the headers of html files
//CSS : css file path
//Title Title of the page
type Header struct {
	CSS   string
	Title string
}

// Footer parameter of the footer of html files
// Js js file path
type Footer struct {
	Js string
}

// HeaderDashboard paramter of the header of the dashboard page
type HeaderDashboard struct {
	OauthLink string
}
