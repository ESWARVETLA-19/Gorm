package models

import "gorm.io/gorm"

// Registration model
type Registration struct {
	gorm.Model
	EventID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
}
