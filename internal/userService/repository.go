package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user *User) (*User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *User) (*User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUserByID(id uint, user *User) (*User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := r.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
