package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyBenefit struct {
	gorm.Model
	RoleTypeIdFrom            string
	RoleTypeIdTo              string
	PartyIdFrom               string
	PartyIdTo                 string
	BenefitTypeId             string
	FromDate                  time.Time
	ThruDate                  time.Time
	PeriodTypeId              string
	Cost                      string
	ActualEmployerPaidPercent float64
	AvailableTime             string

	CreatedAt time.Time
	UpdatedAt time.Time
}
