package models

import (
	"time"
)

type Address struct {
	ID			uint		`gorm:"primary_key"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Address		string		`gorm:"unique"`
	Name		string
	Blocked		bool
	BlockedAt	time.Time
}
