package user

import (
	"gorm.io/gorm"
)

// User модель пользователя
type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string
	Name     string
}
