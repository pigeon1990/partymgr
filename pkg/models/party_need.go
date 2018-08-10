package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyNeed struct {
	gorm.Model
	PartyNeedId          string
	PartyId              string
	RoleTypeId           string
	PartyTypeId          string
	NeedTypeId           string
	CommunicationEventId string
	ProductId            string
	ProductCategoryId    string
	VisitId              string
	DatetimeRecorded     time.Time
	Description          string

	CreatedAt time.Time
	UpdatedAt time.Time
}
