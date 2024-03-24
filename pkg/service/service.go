package service

import "titleNewsRestApi/pkg/repository"

type Authorization interface {
}

type TitleList interface {
}

type Service struct {
	Authorization
	TitleList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
