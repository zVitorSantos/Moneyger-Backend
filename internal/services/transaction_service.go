package services

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"github.com/zVitorSantos/Moneyger-Backend/internal/repositories"
)

type TransactionService struct {
	Repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{Repo: repo}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	return s.Repo.CreateTransaction(transaction)
}

func (s *TransactionService) GetTransactions() ([]models.Transaction, error) {
	return s.Repo.GetTransactions()
}

func (s *TransactionService) GetTransactionByID(id uint) (*models.Transaction, error) {
	return s.Repo.GetTransactionByID(id)
}

func (s *TransactionService) UpdateTransaction(transaction *models.Transaction) error {
	return s.Repo.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	return s.Repo.DeleteTransaction(id)
}
