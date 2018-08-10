package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyClassificationGroup struct {
	gorm.Model
	PartyClassificationGroupId string
	PartyClassificationTypeId  string
	ParentGroupId              string
	Description                string

	CreatedAt time.Time
	UpdatedAt time.Time
}
