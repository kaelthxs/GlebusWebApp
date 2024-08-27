package repository

import (
	"github.com/jmoiron/sqlx"
	"oooGlebusApi"
)

type Authorization interface {
	CreateClient(client oooGlebusApi.Client) (int, error)
	GetClient(username, password string) (oooGlebusApi.Client, error)
}

type Client interface {
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
