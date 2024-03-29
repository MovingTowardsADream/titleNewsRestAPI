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

func (s *TitleListService) GetAll(userId int) ([]news.Title, error) {
	return s.repo.GetAll(userId)
}

func (s *TitleListService) GetById(userId, listId int) (news.Title, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TitleListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
