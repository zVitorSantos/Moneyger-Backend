package repositories

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *TransactionRepository) GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.DB.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r *TransactionRepository) DeleteTransaction(id uint) error {
	return r.DB.Delete(&models.Transaction{}, id).Error
}
