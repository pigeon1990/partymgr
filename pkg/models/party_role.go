package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyRole struct {
	gorm.Model
	PartyId    string
	RoleTypeId string

	CreatedAt time.Time
	UpdatedAt time.Time
}
