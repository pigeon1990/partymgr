package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyRateNew struct {
	gorm.Model
	PartyId        string
	RateTypeId     string
	DefaultRate    string
	PercentageUsed float64
	FromDate       time.Time
	ThruDate       time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
