package service

import (
	"github.com/gin-gonic/gin"
	"oooGlebusApi"
	"oooGlebusApi/pkg/repository"
)

type Authorization interface {
	CreateClient(client oooGlebusApi.Client) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Client interface {
	GetAllClient(c *gin.Context)
	GetClientById(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeleteClient(c *gin.Context)
}

type Album interface {
}

type Music interface {
}

type Service struct {
	Authorization
	Client
	Album
	Music
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
