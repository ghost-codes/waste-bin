package db

import "gorm.io/gorm"

type Store struct {
	*gorm.DB
}
