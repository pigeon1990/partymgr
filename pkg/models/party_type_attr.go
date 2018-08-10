package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyTypeAttr struct {
	gorm.Model
	PartyTypeId string
	AttrName    string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}
