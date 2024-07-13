package repositories

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	return r.DB.Create(account).Error
}

func (r *AccountRepository) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := r.DB.Find(&accounts).Error
	return accounts, err
}

func (r *AccountRepository) GetAccountByID(id uint) (*models.Account, error) {
	var account models.Account
	err := r.DB.First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) UpdateAccount(account *models.Account) error {
	return r.DB.Save(account).Error
}

func (r *AccountRepository) DeleteAccount(id uint) error {
	return r.DB.Delete(&models.Account{}, id).Error
}
