package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyDataSource struct {
	gorm.Model
	PartyId      string
	DataSourceId string
	FromDate     time.Time
	VisitId      string
	Comments     string
	IsCreate     string

	CreatedAt time.Time
	UpdatedAt time.Time
}
