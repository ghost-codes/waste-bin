package db

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type BinStatus string
type BinType string

const (
	Pending   BinStatus = "pending"
	Delivered BinStatus = "delivered"
)

const (
	Large  BinType = "large"
	Medium BinType = "medium"
	Small  BinType = "small"
)

func (ct *BinType) Scan(value interface{}) error {
	*ct = BinType(value.([]byte))
	return nil
}

func (ct BinType) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *BinStatus) Scan(value interface{}) error {
	*ct = BinStatus(value.([]byte))
	return nil
}

func (ct BinStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Bin struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID string
	Status BinStatus `gorm:"type:bin_status;default:pending"`
	Type   BinType   `gorm:"type:bin_type"`
}
