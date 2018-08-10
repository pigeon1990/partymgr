package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyQual struct {
	gorm.Model
	PartyId           string
	PartyQualTypeId   string
	QualificationDesc string
	Title             string
	StatusId          string
	VerifStatusId     string
	FromDate          time.Time
	ThruDate          time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
