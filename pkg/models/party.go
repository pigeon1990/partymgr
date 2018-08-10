package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Party struct {
	gorm.Model
	PartyId                 string
	PartyTypeId             string
	ExternalId              string
	PreferredCurrencyUomId  string
	Description             string
	StatusId                string
	CreatedDate             time.Time
	CreatedByUserLogin      string
	LastModifiedDate        time.Time
	LastModifiedByUserLogin string
	DataSourceId            string
	IsUnread                string

	CreatedAt time.Time
	UpdatedAt time.Time
}
