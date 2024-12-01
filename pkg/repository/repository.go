package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"oooGlebusApi"
)

type Authorization interface {
	CreateClient(client oooGlebusApi.Client) (int, error)
	GetClient(username, password string) (oooGlebusApi.Client, error)
}

type Client interface {
	GetAllClient(c *gin.Context)
	GetClientById(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeleteClient(c *gin.Context)
}

type Album interface {
	createAlbum(c *gin.Context)
	getAllAlbum(c *gin.Context)
	getAllAlbumsByRating(c *gin.Context)
	getAlbumById(c *gin.Context)
	updateAlbum(c *gin.Context)
	deleteAlbum(c *gin.Context)
}

type Music interface {
	createMusic(c *gin.Context)
	getAllMusic(c *gin.Context)
	getMusicById(c *gin.Context)
	updateMusic(c *gin.Context)
	deleteMusic(c *gin.Context)
}

type Review interface {
	createReview(c *gin.Context)
	getAllReview(c *gin.Context)
	getReviewById(c *gin.Context)
	updateReview(c *gin.Context)
	deleteReview(c *gin.Context)
}

type Repository struct {
	Authorization
	Client
	Album
	Music
	Review
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
