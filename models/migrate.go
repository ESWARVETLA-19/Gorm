package models

import (
	"fmt"
	"gorm.io/gorm"
)

// Migrate handles the database migrations
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Event{}, &User{}, &Registration{})
	if err != nil {
		fmt.Printf("Failed to migrate tables: %v\n", err)
		return
	}
	fmt.Println("Database migrated successfully!")
}
