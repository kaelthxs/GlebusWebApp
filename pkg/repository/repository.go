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
}

type Music interface {
}

type Repository struct {
	Authorization
	Client
	Album
	Music
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
