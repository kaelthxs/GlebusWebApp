package oooGlebusApi

type Client struct {
	Id           int    `json:"id" db:"id"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone_Number string `json:"phone_number" binding:"required"`
	Role         string `json:"role" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Image_uri    string `json:"image_uri" binding:"required"`
}
