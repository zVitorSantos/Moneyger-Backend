package repositories

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"gorm.io/gorm"
)

type DebtLoanRepository struct {
	DB *gorm.DB
}

func NewDebtLoanRepository(db *gorm.DB) *DebtLoanRepository {
	return &DebtLoanRepository{DB: db}
}

func (r *DebtLoanRepository) CreateDebtLoan(debtLoan *models.DebtLoan) error {
	return r.DB.Create(debtLoan).Error
}

func (r *DebtLoanRepository) GetDebtLoans() ([]models.DebtLoan, error) {
	var debtLoans []models.DebtLoan
	err := r.DB.Find(&debtLoans).Error
	return debtLoans, err
}

func (r *DebtLoanRepository) GetDebtLoanByID(id uint) (*models.DebtLoan, error) {
	var debtLoan models.DebtLoan
	err := r.DB.First(&debtLoan, id).Error
	if err != nil {
		return nil, err
	}
	return &debtLoan, nil
}

func (r *DebtLoanRepository) UpdateDebtLoan(debtLoan *models.DebtLoan) error {
	return r.DB.Save(debtLoan).Error
}

func (r *DebtLoanRepository) DeleteDebtLoan(id uint) error {
	return r.DB.Delete(&models.DebtLoan{}, id).Error
}
