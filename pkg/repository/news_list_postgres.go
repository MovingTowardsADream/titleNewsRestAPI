package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	news "titleNewsRestApi"
)

type TitleListPostgres struct {
	db *sqlx.DB
}

func NewTitleListPostgres(db *sqlx.DB) *TitleListPostgres {
	return &TitleListPostgres{db: db}
}

func (r *TitleListPostgres) Create(userId int, list news.Title) (int, error) {
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (name, user_id) VALUES ($1, $2) RETURNING id", titleTable)
	row := r.db.QueryRow(createListQuery, list.Name, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
