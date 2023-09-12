package db

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type RequestState string
type RequestType string

func (ct *RequestState) Scan(value interface{}) error {
	*ct = RequestState(value.([]byte))
	return nil
}

func (ct RequestState) Value() (driver.Value, error) {
	return string(ct), nil
}

func (ct *RequestType) Scan(value interface{}) error {
	*ct = RequestType(value.([]byte))
	return nil
}

func (ct RequestType) Value() (driver.Value, error) {
	return string(ct), nil
}

const (
	PENDING RequestState = "pending"
	SUCCESS RequestState = "success"
	FAILED  RequestState = "failed"

	EXTRABIN      RequestType = "extra_bin"
	REPLACEMENT   RequestType = "replacement"
	NEWBIN        RequestType = "new_bin"
	SPECIALPICKUP RequestType = "special_pickup"
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
	State           RequestState `gorm:"type:request_state;default:pending"`
	Type            RequestType  `gorm:"type:request_type"`
}

func (store *Store) CreateRequest(request *Request) error {
	if err := store.Create(request).Error; err != nil {
		return err
	}

	return nil
}
