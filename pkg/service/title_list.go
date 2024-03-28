package service

import (
	news "titleNewsRestApi"
	"titleNewsRestApi/pkg/repository"
)

type TitleListService struct {
	repo repository.TitleList
}

func NewTitleListService(repo repository.TitleList) *TitleListService {
	return &TitleListService{repo: repo}
}

func (s *TitleListService) Create(userId int, list news.Title) (int, error) {
	return s.repo.Create(userId, list)
}
