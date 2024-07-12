package models

import (
    "time"
)

type Category struct {
    CategoryID   uint      `gorm:"primaryKey"`
    CategoryName string    `gorm:"unique;not null"`
    CategoryType string    `gorm:"not null"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
