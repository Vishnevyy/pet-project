package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
