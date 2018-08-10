package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyInvitation struct {
	gorm.Model
	PartyInvitationId string
	PartyIdFrom       string
	PartyId           string
	ToName            string
	EmailAddress      string
	StatusId          string
	LastInviteDate    time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
