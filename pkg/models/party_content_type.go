package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyContentType struct {
	gorm.Model
	PartyContentTypeId string
	ParentTypeId       string
	Description        string

	CreatedAt time.Time
	UpdatedAt time.Time
}
