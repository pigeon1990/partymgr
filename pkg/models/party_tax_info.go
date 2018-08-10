package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyTaxInfo struct {
	gorm.Model
	PartyId    string
	GeoId      string
	FromDate   time.Time
	ThruDate   time.Time
	PartyTaxId string
	IsExempt   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
