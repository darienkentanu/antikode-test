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

type GormOutletModel struct {
	db *gorm.DB
}

func NewOutletModel(db *gorm.DB) *GormOutletModel {
	return &GormOutletModel{db: db}
}

type OutletModel interface {
	Insert(outlet Outlet) (Outlet, error)
	GetAll() ([]Outlet, error)
	// Edit(id int, outlet Outlet) (Outlet, error)
	// Delete(id int) (Outlet, error)
}

func (om *GormOutletModel) Insert(outlet Outlet) (Outlet, error) {
	tx := om.db.Begin()
	if err := tx.Create(&outlet).Error; err != nil {
		tx.Rollback()
		return Outlet{}, err
	}
	tx.Commit()
	return outlet, nil
}

func (om *GormOutletModel) GetAll() ([]Outlet, error) {
	var outlets []Outlet
	if err := om.db.Find(&outlets).Error; err != nil {
		return nil, err
	}
	return outlets, nil
}
