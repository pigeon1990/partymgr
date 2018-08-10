package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyNameHistory struct {
	gorm.Model
	PartyId       string
	ChangeDate    time.Time
	GroupName     string
	FirstName     string
	MiddleName    string
	LastName      string
	PersonalTitle string
	Suffix        string

	CreatedAt time.Time
	UpdatedAt time.Time
}
