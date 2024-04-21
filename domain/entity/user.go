package entity

import (
	"gorm.io/gorm"
)

type User struct {
	Base `gorm:"embedded"`
	Name string
	Todo []Todo
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
