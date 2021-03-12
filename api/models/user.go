package models

type User struct {
	ID    int64  `json:"id" gorm:"primarykey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
