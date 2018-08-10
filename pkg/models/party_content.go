package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyContent struct {
	gorm.Model
	PartyId            string
	ContentId          string
	PartyContentTypeId string
	FromDate           time.Time
	ThruDate           time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
