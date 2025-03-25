package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the SQLite database connection
func ConnectDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	fmt.Println("SQLite database connected successfully!")
	return DB, nil
}
