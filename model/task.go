package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Owner  uint `gorm:"not null"`
	UserID uint
	TeamID uint
	User   *User `gorm:"ForeignKey:UserID"`
	Team   *Team `gorm:"ForeignKey:TeamID"`

	Status      int    `gorm:"default:0"`
	Name        string `gorm:"index;not null"`
	Description string
	StartTime   int64
	DueDate     int64
}
