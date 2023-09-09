package db

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID string
	User   UserDetails `gorm:"foreignKey:UserID"`
	Amount int
	Date   time.Time
}
