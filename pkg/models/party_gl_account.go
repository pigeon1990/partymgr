package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyGlAccount struct {
	gorm.Model
	OrganizationPartyId string
	PartyId             string
	RoleTypeId          string
	GlAccountTypeId     string
	GlAccountId         string

	CreatedAt time.Time
	UpdatedAt time.Time
}
