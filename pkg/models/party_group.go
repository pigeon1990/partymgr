package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyGroup struct {
	gorm.Model
	PartyId        string
	GroupName      string
	GroupNameLocal string
	OfficeSiteName string
	AnnualRevenue  string
	NumEmployees   string
	TickerSymbol   string
	Comments       string
	LogoImageUrl   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
