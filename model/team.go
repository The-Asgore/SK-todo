package model

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	TeamName string  `gorm:"unique"`
	User     []*User `gorm:"many2many:user_team;"`
}
