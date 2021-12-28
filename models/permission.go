package models

import (
	"time"
)

type Permission struct {
	ID			uint		`gorm:"primary_key"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	AddressID	uint		`gorm:"foreignKey:AddressID"`
	Allowed		bool		`gorm:"default:false"`
	DecidedAt	time.Time
}
