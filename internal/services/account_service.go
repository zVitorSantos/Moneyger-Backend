package services

import (
	"github.com/zVitorSantos/Moneyger-Backend/internal/models"
	"github.com/zVitorSantos/Moneyger-Backend/internal/repositories"
)

type AccountService struct {
	Repo *repositories.AccountRepository
}

func NewAccountService(repo *repositories.AccountRepository) *AccountService {
	return &AccountService{Repo: repo}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.Repo.CreateAccount(account)
}

func (s *AccountService) GetAccounts() ([]models.Account, error) {
	return s.Repo.GetAccounts()
}

func (s *AccountService) GetAccountByID(id uint) (*models.Account, error) {
	return s.Repo.GetAccountByID(id)
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.Repo.UpdateAccount(account)
}

func (s *AccountService) DeleteAccount(id uint) error {
	return s.Repo.DeleteAccount(id)
}
