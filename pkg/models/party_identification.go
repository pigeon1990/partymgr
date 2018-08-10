package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyIdentification struct {
	gorm.Model
	PartyId                   string
	PartyIdentificationTypeId string
	IdValue                   string

	CreatedAt time.Time
	UpdatedAt time.Time
}
