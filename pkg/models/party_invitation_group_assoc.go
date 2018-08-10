package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyInvitationGroupAssoc struct {
	gorm.Model
	PartyInvitationId string
	PartyIdTo         string

	CreatedAt time.Time
	UpdatedAt time.Time
}
