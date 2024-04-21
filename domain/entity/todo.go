package entity

import (
	"github.com/google/uuid"
)

type Todo struct {
	Base   `gorm:"embedded"`
	Title  string
	Done   bool
	UserID uuid.UUID
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}

func NewTodo(title string, done bool, userID uuid.UUID) *Todo {
	return &Todo{
		Title:  title,
		Done:   done,
		UserID: userID,
	}
}
