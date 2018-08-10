package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyGeoPoint struct {
	gorm.Model
	PartyId    string
	GeoPointId string
	FromDate   time.Time
	ThruDate   time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
