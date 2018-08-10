package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyFixedAssetAssignment struct {
	gorm.Model
	PartyId       string
	RoleTypeId    string
	FixedAssetId  string
	FromDate      time.Time
	ThruDate      time.Time
	AllocatedDate time.Time
	StatusId      string
	Comments      string

	CreatedAt time.Time
	UpdatedAt time.Time
}
