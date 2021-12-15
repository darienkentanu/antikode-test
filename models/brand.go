package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name     string `gorm:"type:varchar(200);not null;unique" json:"name"`
	Logo     string `gorm:"type:longtext" json:"logo"`
	Banner   string `gorm:"type:longtext" json:"banner"`
	Outlets  []Outlet
	Products []Product
}
