package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyContactMechPurpose struct {
	gorm.Model
	PartyId                  string
	ContactMechId            string
	ContactMechPurposeTypeId string
	FromDate                 time.Time
	ThruDate                 time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
