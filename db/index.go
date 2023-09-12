package db

import (
	"fmt"

	models "ghost-codes/waste-bin/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Migrations struct {
	DB     *gorm.DB
	Models []interface{}
}

func RunMigrations(migrations Migrations) error {
	for _, model := range migrations.Models {
		err := migrations.DB.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("Error running migrations, %s\nModel %v", err, model)
		}
	}
	return nil
}

func NewGorm(dbConnLink string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbConnLink), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("DB could not be initialized: %s", err)
	}

	migrations := Migrations{
		DB: db,
		Models: []interface{}{
			&models.UserDetails{},
			&models.Payment{},
			&models.SpecialPickup{},
			&models.Request{},
		},
	}

	err = RunMigrations(migrations)
	if err != nil {
		return nil, err
	}

	return db, nil
}
