package db

import (
	"time"

	"gorm.io/gorm"
)

type SpecialPickup struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	UserID       string
	User         UserDetails `gorm:"foreignKey:UserID"`
	DateOfPickup time.Time
}
