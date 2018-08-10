package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyQualType struct {
	gorm.Model
	PartyQualTypeId string
	ParentTypeId    string
	HasTable        string
	Description     string

	CreatedAt time.Time
	UpdatedAt time.Time
}
