package userService

import (
	"pet-project/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"-"`
	Tasks    []taskService.Task `json:"tasks" gorm:"foreignKey:UserID"`
}
