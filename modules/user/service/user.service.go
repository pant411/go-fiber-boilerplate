package service

import (
	"go-fiber-boilerplate/modules/user/model"
	"go-fiber-boilerplate/modules/user/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// FindAllUser returns all users
func (s *UserService) FindAllUser() ([]model.User, error) {
	return s.userRepo.FindAllUser()
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}

// FindOneUser returns a user by ID
func (s *UserService) FindOneUser(id uint) (*model.User, error) {
	return s.userRepo.FindOneUser(id)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id uint, user *model.User) error {
	return s.userRepo.UpdateUser(id, user)
}
