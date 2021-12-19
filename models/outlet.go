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
	Edit(id int, outlet Outlet) (Outlet, error)
	Delete(id int) (Outlet, error)
	GetOutletById(id int) (Outlet, error)
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
	tx := om.db.Begin()
	if err := tx.Find(&outlets).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return outlets, nil
}

func (om *GormOutletModel) GetOutletById(id int) (Outlet, error) {
	var outlet Outlet
	tx := om.db.Begin()
	if err := tx.First(&outlet, id).Error; err != nil {
		tx.Rollback()
		return Outlet{}, err
	}
	tx.Commit()
	return outlet, nil
}

func (om *GormOutletModel) Edit(id int, outlet Outlet) (Outlet, error) {
	var outletOutput Outlet
	tx := om.db.Begin()
	if err := tx.Model(&outletOutput).Where("id=?", id).Updates(Outlet{
		Name:      outlet.Name,
		Picture:   outlet.Picture,
		Address:   outlet.Address,
		Longitude: outlet.Longitude,
		Latitude:  outlet.Latitude,
		Distance:  outlet.Distance,
	}).Error; err != nil {
		tx.Rollback()
		return Outlet{}, err
	}
	tx.Commit()
	return outletOutput, nil
}

func (om *GormOutletModel) Delete(id int) (Outlet, error) {
	var outlet Outlet
	tx := om.db.Begin()
	if err := tx.First(&outlet, id).Error; err != nil {
		tx.Rollback()
		return Outlet{}, err
	}
	if err := tx.Delete(&Outlet{}, id).Error; err != nil {
		tx.Rollback()
		return Outlet{}, err
	}
	tx.Commit()
	return outlet, nil
}
