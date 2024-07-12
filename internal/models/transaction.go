package models

import (
    "time"
)

type Transaction struct {
    TransactionID   uint      `gorm:"primaryKey"`
    UserID          uint      `gorm:"not null"`
    AccountID       uint      `gorm:"not null"`
    Amount          float64   `gorm:"type:numeric(15,2);not null"`
    TransactionType string    `gorm:"not null"`
    CategoryID      uint
    Description     string
    TransactionDate time.Time `gorm:"type:date;not null"`
    ReceiptURL      string    `gorm:"size:255"`
    CreatedAt       time.Time `gorm:"autoCreateTime"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime"`
    
    User            User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    Account         Account   `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
    Category        Category  `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL"`
}
