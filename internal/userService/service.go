package userService

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user *User) (*User, error)
	DeleteUserByID(id uint) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *User) (*User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUserByID(id uint, user *User) (*User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *userService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
