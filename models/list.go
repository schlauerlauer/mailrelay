package models

import (
	"time"
)

type List struct {
	ID			uint 		`gorm:"primary_key"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Addresses	[]Address	`gorm:"many2many:list_addresses"`
	Name		string
	Address		Address
	AddressID	uint		`gorm:"foreignKey:AddressID"`
	Domains		[]Domain	`gorm:"many2many:list_domains"`
	ListType	ListType
}

type ListType int

const (
	Parents ListType = iota
	Leaders
)

var lt = [...]string{
	"Parents",
	"Leaders",
}

func (t ListType) String() string {
	return lt[t]
}

func (t ListType) EnumIndex() int {
	return int(t)
}
