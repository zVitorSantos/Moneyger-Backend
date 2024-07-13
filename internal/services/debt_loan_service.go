package services

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"github.com/zVitorSantos/Moneyger-Backend/internal/repositories"
)

type DebtLoanService struct {
	Repo *repositories.DebtLoanRepository
}

func NewDebtLoanService(repo *repositories.DebtLoanRepository) *DebtLoanService {
	return &DebtLoanService{Repo: repo}
}

func (s *DebtLoanService) CreateDebtLoan(debtLoan *models.DebtLoan) error {
	return s.Repo.CreateDebtLoan(debtLoan)
}

func (s *DebtLoanService) GetDebtLoans() ([]models.DebtLoan, error) {
	return s.Repo.GetDebtLoans()
}

func (s *DebtLoanService) GetDebtLoanByID(id uint) (*models.DebtLoan, error) {
	return s.Repo.GetDebtLoanByID(id)
}

func (s *DebtLoanService) UpdateDebtLoan(debtLoan *models.DebtLoan) error {
	return s.Repo.UpdateDebtLoan(debtLoan)
}

func (s *DebtLoanService) DeleteDebtLoan(id uint) error {
	return s.Repo.DeleteDebtLoan(id)
}
