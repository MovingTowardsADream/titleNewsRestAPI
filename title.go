package news

type Title struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name" binding:"required"`
	UserId int    `json:"user_id" db:"user_id"`
}
