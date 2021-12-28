package models

import (
	"time"
)

type Mail struct {
	ID			uint		`gorm:"primary_key"`
	CreatedAt	time.Time
	FromID		uint		`gorm:"foreignKey:AddressID"`
	From		Address
	ToId		uint		`gorm:"foreignKey:AddressId"`
	To			Address
	Allowed		bool
}
