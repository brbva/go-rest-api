package service

import (
	"gorest/internal/user/model"
	"gorest/internal/user/repository"
)

// UserService defines methods for handling user-related logic
type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	CreateUser(user model.User) (int, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id int) (model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) CreateUser(user model.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) UpdateUser(user model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
