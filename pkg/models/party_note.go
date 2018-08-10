package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyNote struct {
	gorm.Model
	PartyId string
	NoteId  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
