package db

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type RequestState string

func (ct *RequestState) Scan(value interface{}) error {
	*ct = RequestState(value.([]byte))
	return nil
}

func (ct RequestState) Value() (driver.Value, error) {
	return string(ct), nil
}

const (
	PENDING RequestState = "pending"
	SUCCESS RequestState = "success"
	FAILED  RequestState = "failed"
)

type Request struct {
	gorm.Model
	ID              uint
	UserID          string
	User            UserDetails `gorm:"foreignKey:UserID"`
	PaymentID       *uint
	Payment         *Payment
	DateOfDropOff   time.Time
	BinID           *uint
	Bin             *Bin
	SpecialPickupID *uint
	SpecialPickup   *SpecialPickup
	State           RequestState `gorm:"type:request_state"`
}
