package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyIdentificationType struct {
	gorm.Model
	PartyIdentificationTypeId string
	ParentTypeId              string
	HasTable                  string
	Description               string

	CreatedAt time.Time
	UpdatedAt time.Time
}
