package oooGlebusApi

type Album struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Countofplays int    `json:"countofplays"`
	Rating       int    `json:"rating"`
	Countofmusic int    `json:"countofmusic"`
	Status       string `json:"status"`
	Image_uri    string `json:"image_uri"`
}
