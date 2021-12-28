package models

import (
	"time"
)

type Domain struct {
	ID			uint		`gorm:"primary_key"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Domain		string		`gorm:"unique"`
}
