package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID           int    `json:"id", gorm:"PrimaryKey"`
	Name         string `json:"name"`
	Login        string `json:"login"`
	PasswordHash string `json:"passwordHash"`
}
