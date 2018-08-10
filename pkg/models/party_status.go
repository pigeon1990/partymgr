package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyStatus struct {
	gorm.Model
	StatusId            string
	PartyId             string
	StatusDate          time.Time
	ChangeByUserLoginId string

	CreatedAt time.Time
	UpdatedAt time.Time
}
