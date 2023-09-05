package db

import (
	"time"

	"database/sql/driver"

	"gorm.io/gorm"
)

type TodoType string

const (
	MONTHLY TodoType = "MONTHLY"
	DAILY   TodoType = "DAILY"
)

func (ct *TodoType) Scan(value interface{}) error {
	*ct = TodoType(value.([]byte))
	return nil
}

func (ct TodoType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Todo struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Note      string
	UID       string
	Type      TodoType `gorm:"type:todo_type"`
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (store *Store) GetUserTodos(uid string) ([]Todo, error) {
	var todos []Todo
	if err := store.Model(&Todo{UID: uid}).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (store *Store) CreateTodo(todo *Todo) error {
	if err := store.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

func (store *Store) DeleteTodo(id uint) error {
	if err := store.Delete(&Todo{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func (store *Store) UpdateTodo(todo *Todo) error {
	if err := store.Save(todo).Error; err != nil {
		return err
	}
	return nil
}
