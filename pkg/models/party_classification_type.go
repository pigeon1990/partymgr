package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyClassificationType struct {
	gorm.Model
	PartyClassificationTypeId string
	ParentTypeId              string
	HasTable                  string
	Description               string

	CreatedAt time.Time
	UpdatedAt time.Time
}
