package taskService

type TaskService interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTask() ([]Task, error)
	UpdateTaskByID(id uint, task *Task) (*Task, error)
	DeleteTaskByID(id uint) error
	GetTasksForUser(userID uint) ([]Task, error)
	GetTaskByID(taskID uint) (*Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *Task) (*Task, error) {
	return s.repo.CreateTask(task)
}

func (s *taskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func (s *taskService) UpdateTaskByID(id uint, task *Task) (*Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *taskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *taskService) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksForUser(userID)
}

func (s *taskService) GetTaskByID(taskID uint) (*Task, error) {
	return s.repo.GetTaskByID(taskID)
}
