package entiity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primary_key"`
	Username  string `gorm:"size:50;unique;uniqueIndex"`
	Password  string `gorm:"size:255"`
	Name      string `gorm:"size:100"`
	Role      string `gorm:"size:15"`
	Avatar    string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
