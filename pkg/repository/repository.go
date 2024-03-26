package repository

import (
	"github.com/jmoiron/sqlx"
	news "titleNewsRestApi"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
}

type TitleList interface {
}

type Repository struct {
	Authorization
	TitleList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
