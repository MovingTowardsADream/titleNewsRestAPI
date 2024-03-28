package news

type Title struct {
	Id     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	UserId int    `json:"user_id"`
}
