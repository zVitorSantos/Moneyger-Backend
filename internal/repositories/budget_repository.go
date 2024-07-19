package repositories

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"gorm.io/gorm"
)

type BudgetRepository struct {
	DB *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{DB: db}
}

func (r *BudgetRepository) CreateBudget(budget *models.Budget) error {
	return r.DB.Create(budget).Error
}

func (r *BudgetRepository) GetBudgets() ([]models.Budget, error) {
	var budgets []models.Budget
	err := r.DB.Find(&budgets).Error
	return budgets, err
}

func (r *BudgetRepository) GetBudgetByID(id uint) (*models.Budget, error) {
	var budget models.Budget
	err := r.DB.First(&budget, id).Error
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepository) UpdateBudget(budget *models.Budget) error {
	return r.DB.Save(budget).Error
}

func (r *BudgetRepository) DeleteBudget(id uint) error {
	return r.DB.Delete(&models.Budget{}, id).Error
}
