package models

import (
	"time"
)

type Budget struct {
	BudgetID     uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"not null"`
	CategoryID   uint      `gorm:"not null"`
	BudgetAmount float64   `gorm:"type:numeric(15,2);not null"`
	BudgetMonth  time.Time `gorm:"type:date;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	User     User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Category Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}
