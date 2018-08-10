package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartySkill struct {
	gorm.Model
	PartyId          string
	SkillTypeId      string
	YearsExperience  string
	Rating           string
	SkillLevel       string
	StartedUsingDate time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
