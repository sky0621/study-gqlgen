package model

type Todo struct {
	ID     int64  `json:"id"`
	Task   string `json:"task"`
	UserID int64  `json:"user_id" db:"user_id"`
}
