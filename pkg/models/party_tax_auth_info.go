package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyTaxAuthInfo struct {
	gorm.Model
	PartyId        string
	TaxAuthGeoId   string
	TaxAuthPartyId string
	FromDate       time.Time
	ThruDate       time.Time
	PartyTaxId     string
	IsExempt       string
	IsNexus        string

	CreatedAt time.Time
	UpdatedAt time.Time
}
