package models

import (
	"errors"
	"gorm.io/gorm"
	"myproject/db"
	"time"
)

//Event model using GORM
type Event struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Location    string    `gorm:"not null" json:"location"`
	DateTime    time.Time `gorm:"not null" json:"date_time"`
	UserID      uint      `gorm:"not null" json:"user_id"`
}

// Save saves the event into the database
func (e *Event) Save() error {
	result := db.DB.Create(&e)
	return result.Error
}

// GetAllEvents retrieves all events from the database
func GetAllEvents() ([]Event, error) {
	var events []Event
	result := db.DB.Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

// GetEvent retrieves a single event by ID
func GetEvent(id uint) (*Event, error) {
	var event Event
	result := db.DB.First(&event, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, result.Error
	}
	return &event, nil
}

// Update modifies the existing event
func (e *Event) Update() error {
	result := db.DB.Save(&e)
	return result.Error
}

// Delete removes the event from the database
func (e *Event) Delete() error {
	result := db.DB.Delete(&e)
	return result.Error
}

// Register adds a registration for the event
func (e *Event) Register(userID uint) error {
	registration := Registration{
		EventID: e.ID,
		UserID:  userID,
	}
	result := db.DB.Create(&registration)
	return result.Error
}

// Unregister removes a registration for the event
func (e *Event) Unregister(userID uint) error {
	result := db.DB.Where("event_id = ? AND user_id = ?", e.ID, userID).Delete(&Registration{})
	return result.Error
}
