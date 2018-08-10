package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyRelationship struct {
	gorm.Model
	PartyIdFrom             string
	PartyIdTo               string
	RoleTypeIdFrom          string
	RoleTypeIdTo            string
	FromDate                time.Time
	ThruDate                time.Time
	StatusId                string
	RelationshipName        string
	SecurityGroupId         string
	PriorityTypeId          string
	PartyRelationshipTypeId string
	PermissionsEnumId       string
	PositionTitle           string
	Comments                string

	CreatedAt time.Time
	UpdatedAt time.Time
}
