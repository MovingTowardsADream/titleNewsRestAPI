package service

import (
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TitleList interface {
}

type Service struct {
	Authorization
	TitleList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthServices(repos.Authorization),
	}
}
