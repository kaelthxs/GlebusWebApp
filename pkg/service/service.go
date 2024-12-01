package service

import (
	"github.com/gin-gonic/gin"
	"oooGlebusApi"
	"oooGlebusApi/pkg/repository"
)

type Authorization interface {
	CreateClient(client oooGlebusApi.Client) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	Client
	Album
	Music
	Review
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
