package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyRelationshipType struct {
	gorm.Model
	PartyRelationshipTypeId string
	ParentTypeId            string
	HasTable                string
	PartyRelationshipName   string
	Description             string
	RoleTypeIdValidFrom     string
	RoleTypeIdValidTo       string

	CreatedAt time.Time
	UpdatedAt time.Time
}
