package services

import (
	"go-chi-api/model"
	"go-chi-api/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	//return &userServiceImpl{userRepository: userRepository}
	return nil
}

func (s *userServiceImpl) CreateUser(user *model.User) error {
	// Implement user creation logic using repository
	return nil
}

// Implement other service methods for user operations
