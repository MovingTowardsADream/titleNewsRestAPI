package repository

import (
	"github.com/jmoiron/sqlx"
	news "titleNewsRestApi"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
	GetUser(username, password string) (news.User, error)
}

type TitleList interface {
	Create(userId int, list news.Title) (int, error)
	GetAll(userId int) ([]news.Title, error)
	GetById(userId, listId int) (news.Title, error)
	Delete(userId, listId int) error
}

type Repository struct {
	Authorization
	TitleList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TitleList:     NewTitleListPostgres(db),
	}
}
