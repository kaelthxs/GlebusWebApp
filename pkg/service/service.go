package service

import (
	"oooGlebusApi"
	"oooGlebusApi/pkg/repository"
)

type Authorization interface {
	CreateClient(client oooGlebusApi.Client) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Client interface {
}

type Album interface {
}

type Music interface {
}

type Review interface {
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
