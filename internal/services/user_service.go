package services

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"github.com/zVitorSantos/Moneyger-Backend/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
