package service

import (
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user news.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TitleList interface {
	Create(userId int, list news.Title) (int, error)
	GetAll(userId int) ([]news.Title, error)
	GetById(userId, listId int) (news.Title, error)
	Delete(userId, listId int) error
}

type Service struct {
	Authorization
	TitleList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthServices(repos.Authorization),
		TitleList:     NewTitleListService(repos.TitleList),
	}
}
