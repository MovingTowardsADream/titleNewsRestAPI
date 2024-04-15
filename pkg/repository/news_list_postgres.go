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

func (r *TitleListPostgres) GetAll(userId int) ([]news.Title, error) {
	var list []news.Title
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", titleTable)
	err := r.db.Select(&list, query, userId)

	return list, err
}

func (r *TitleListPostgres) GetById(userId, listId int) (news.Title, error) {
	var list news.Title
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND user_id=$2", titleTable)
	err := r.db.Get(&list, query, listId, userId)

	return list, err
}

func (r *TitleListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", titleTable)

	_, err := r.db.Exec(query, listId, userId)

	return err
}
