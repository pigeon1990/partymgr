package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyAttribute struct {
	gorm.Model
	PartyId         string
	AttrName        string
	AttrValue       string
	AttrDescription string

	CreatedAt time.Time
	UpdatedAt time.Time
}
