package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyType struct {
	gorm.Model
	PartyTypeId  string
	ParentTypeId string
	HasTable     string
	Description  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
