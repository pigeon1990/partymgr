package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyContactMech struct {
	gorm.Model
	PartyId               string
	ContactMechId         string
	FromDate              time.Time
	ThruDate              time.Time
	RoleTypeId            string
	AllowSolicitation     string
	Extension             string
	Verified              string
	Comments              string
	YearsWithContactMech  string
	MonthsWithContactMech string

	CreatedAt time.Time
	UpdatedAt time.Time
}
