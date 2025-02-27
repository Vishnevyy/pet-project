package taskService

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTask() ([]Task, error)
	UpdateTaskByID(id uint, task *Task) (*Task, error)
	DeleteTaskByID(id uint) error
	GetTasksForUser(userID uint) ([]Task, error)
	GetTaskByID(taskID uint) (*Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task *Task) (*Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) GetAllTask() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTaskByID(id uint, task *Task) (*Task, error) {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return nil, err
	}

	existingTask.Title = task.Title
	existingTask.Completed = task.Completed

	if err := r.db.Save(&existingTask).Error; err != nil {
		return nil, err
	}

	return &existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetTasksForUser(userID uint) ([]Task, error) {
	var tasks []Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetTaskByID(taskID uint) (*Task, error) {
	var task Task
	if err := r.db.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
