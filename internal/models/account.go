package models

import (
    "time"
)

type Account struct {
    AccountID   uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"not null"`
    AccountName string    `gorm:"not null"`
    AccountType string    `gorm:"not null"` // Example: bank, card, wallet...
    Balance     float64   `gorm:"type:numeric(15,2);default:0.0"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`
    
    User        User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
