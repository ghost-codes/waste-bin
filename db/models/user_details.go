package db

import (
	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	UserID     string `gorm:"primaryKey"`
	Fullname   string
	LocationID int `gorm:"foreignKey"`
	Location   Location
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Location struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Lat  float32
	Long float32
	Name string
}

func (store *Store) FetchUserByUID(uid string) (*UserDetails, error) {
	user := &UserDetails{UserID: uid}
	if err := store.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (store *Store) CreateUser(user *UserDetails) error {
	if err := store.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (store *Store) UpdateUser(user *UserDetails) error {
	if err := store.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (store *Store) DeleteUserByUID(uid string) error {
	if err := store.Delete(&UserDetails{UserID: uid}).Error; err != nil {
		return err
	}
	return nil
}
