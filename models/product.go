package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `gorm:"type:varchar(200);unique;not null" json:"name"`
	Picture string  `gorm:"type:longtext" json:"picture"`
	Price   float32 `gorm:"type:decimal(20,2);not null" json:"price"`
	BrandID uint    `gorm:"type:bigint;not null" json:"brand_id"`
}
