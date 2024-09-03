package oooGlebusApi

type Review struct {
	Id          int    `json:"id"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Client_id   int    `json:"client_id"`
	Music_id    int    `json:"music_id"`
}
