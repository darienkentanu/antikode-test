package models

import "gorm.io/gorm"

type Outlet struct {
	gorm.Model
	Name      string  `gorm:"type:varchar(200);not null; unique" json:"name"`
	Picture   string  `gorm:"type:longtext" json:"picture"`
	Address   string  `gorm:"type:longtext;not null" json:"address"`
	Longitude string  `gorm:"type:longtext" json:"longitute"`
	Latitude  string  `gorm:"type:longtext" json:"latitude"`
	BrandID   uint    `gorm:"type:bigint;not null" json:"brand_id"`
	Distance  float32 `gorm:"type:decimal(10,2)" json:"-"`
}
