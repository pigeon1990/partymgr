package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyRate struct {
	gorm.Model
	PartyId       string
	RateTypeId    string
	CurrencyUomId string
	DefaultRate   string
	FromDate      time.Time
	ThruDate      time.Time
	Rate          string

	CreatedAt time.Time
	UpdatedAt time.Time
}
