package services

import (
    "github.com/ashishthakur913/CaterEase/models"
    "github.com/ashishthakur913/CaterEase/repositories"
)

type UserService struct {
    userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.userRepository.CreateUser(user)
}