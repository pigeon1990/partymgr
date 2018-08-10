package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyInvitationRoleAssoc struct {
	gorm.Model
	PartyInvitationId string
	RoleTypeId        string

	CreatedAt time.Time
	UpdatedAt time.Time
}
