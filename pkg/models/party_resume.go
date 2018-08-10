package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyResume struct {
	gorm.Model
	ResumeId   string
	PartyId    string
	ContentId  string
	ResumeDate time.Time
	ResumeText string

	CreatedAt time.Time
	UpdatedAt time.Time
}
