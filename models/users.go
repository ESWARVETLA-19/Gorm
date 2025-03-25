package models

import (
	"errors"
	"gorm.io/gorm"
	 "myproject/db"
	"myproject/utils"
)

// User struct (GORM model)
type User struct {
	gorm.Model
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Events   []Event   `gorm:"foreignKey:UserID"`   // One-to-Many relationship
}

// Save creates a new user with hashed password
func (user *User) Save() error {
	// Hash the password before saving
	hashedPassword, err := utils.HashGenerator(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save the user using GORM
	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ValidateCreds checks if the email and password are valid
func (user *User) ValidateCreds() error {
	var dbUser User

	// Fetch user by email
	result := db.DB.Where("email = ?", user.Email).First(&dbUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("email or password is incorrect")
		}
		return result.Error
	}

	// Verify the password
	passwdValid := utils.CheckHashPassword(user.Password, dbUser.Password)
	if !passwdValid {
		return errors.New("email or password is incorrect")
	}

	// Set the authenticated user's ID
	user.ID = dbUser.ID
	return nil
}
