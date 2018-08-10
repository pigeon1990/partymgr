package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyCarrierAccount struct {
	gorm.Model
	PartyId        string
	CarrierPartyId string
	FromDate       time.Time
	ThruDate       time.Time
	AccountNumber  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
