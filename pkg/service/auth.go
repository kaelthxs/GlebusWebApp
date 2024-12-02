package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"oooGlebusApi"
	"oooGlebusApi/pkg/repository"
	"time"
)

const (
	salt       = "sdfihuf77dshdssd"
	signingKey = "djfsdhuuteutrnfj3434354"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	ClientId int `json:"client_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) ParseToken(token string) (int, error) {
	panic("implement me")
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateClient(client oooGlebusApi.Client) (int, error) {
	client.Password = generatePasswordHash(client.Password)
	return s.repo.CreateClient(client)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	client, err := s.repo.GetClient(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		client.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
