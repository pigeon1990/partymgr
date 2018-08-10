package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyIcsAvsOverride struct {
	gorm.Model
	PartyId          string
	AvsDeclineString string

	CreatedAt time.Time
	UpdatedAt time.Time
}
