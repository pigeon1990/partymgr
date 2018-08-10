package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PartyProfileDefault struct {
	gorm.Model
	PartyId         string
	ProductStoreId  string
	DefaultShipAddr string
	DefaultBillAddr string
	DefaultPayMeth  string
	DefaultShipMeth string

	CreatedAt time.Time
	UpdatedAt time.Time
}
