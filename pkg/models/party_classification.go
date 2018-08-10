package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyClassification struct {
	gorm.Model
	PartyId                    string
	PartyClassificationGroupId string
	FromDate                   time.Time
	ThruDate                   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
