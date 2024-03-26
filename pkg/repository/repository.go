package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TitleList interface {
}

type Repository struct {
	Authorization
	TitleList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
