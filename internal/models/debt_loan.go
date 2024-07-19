package models

import (
	"time"
)

type DebtLoan struct {
	DebtLoanID   uint    `gorm:"primaryKey"`
	UserID       uint    `gorm:"not null"`
	Amount       float64 `gorm:"type:numeric(15,2);not null"`
	DebtLoanType string  `gorm:"not null"`
	Description  string
	DueDate      time.Time `gorm:"type:date"`
	PaymentAlert bool      `gorm:"default:false"`
	Paid         bool      `gorm:"default:false"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
