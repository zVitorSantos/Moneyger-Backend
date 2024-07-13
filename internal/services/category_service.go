package services

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"github.com/zVitorSantos/Moneyger-Backend/internal/repositories"
)

type CategoryService struct {
	Repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.Repo.CreateCategory(category)
}

func (s *CategoryService) GetCategories() ([]models.Category, error) {
	return s.Repo.GetCategories()
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.Repo.GetCategoryByID(id)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return s.Repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.Repo.DeleteCategory(id)
}
