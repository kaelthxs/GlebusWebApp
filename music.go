package oooGlebusApi

type Music struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Rating       int    `json:"rating"`
	Countofplays int    `json:"countofplays"`
	Album_id     int    `json:"album_id"`
	Image_uri    string `json:"image_uri"`
	Audio_uri    string `json:"audio_uri"`
}
