package repository

type Authorization interface {
}

type TitleList interface {
}

type Repository struct {
	Authorization
	TitleList
}

func NewRepository() *Repository {
	return &Repository{}
}
