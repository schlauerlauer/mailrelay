package models

import (
	"time"
)

type Mail struct {
	ID			uint		`gorm:"primary_key"`
	CreatedAt	time.Time
	FromID		uint		`gorm:"foreignKey:AddressID"`
	From		Address
	To			[]Address	`gorm:"many2many:mail_addresses"`
	Allowed		bool
}
