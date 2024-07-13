package services

import (
    "github.com/zVitorSantos/Moneyger-Backend/internal/models"
    "github.com/zVitorSantos/Moneyger-Backend/internal/repositories"
)

type BudgetService struct {
    Repo *repositories.BudgetRepository
}

func NewBudgetService(repo *repositories.BudgetRepository) *BudgetService {
    return &BudgetService{Repo: repo}
}

func (s *BudgetService) CreateBudget(budget *models.Budget) error {
    return s.Repo.CreateBudget(budget)
}

func (s *BudgetService) GetBudgets() ([]models.Budget, error) {
    return s.Repo.GetBudgets()
}

func (s *BudgetService) GetBudgetByID(id uint) (*models.Budget, error) {
    return s.Repo.GetBudgetByID(id)
}

func (s *BudgetService) UpdateBudget(budget *models.Budget) error {
    return s.Repo.UpdateBudget(budget)
}

func (s *BudgetService) DeleteBudget(id uint) error {
    return s.Repo.DeleteBudget(id)
}
