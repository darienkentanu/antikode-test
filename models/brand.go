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

type GormBrandModel struct {
	db *gorm.DB
}

func NewBrandModel(db *gorm.DB) *GormBrandModel {
	return &GormBrandModel{db: db}
}

type BrandModel interface {
	GetAll() ([]Brand, error)
	Insert(Brand) (Brand, error)
	Edit(id int, brand Brand) (Brand, error)
	Delete(id int) (Brand, error)
}

func (bm *GormBrandModel) GetAll() ([]Brand, error) {
	var brand []Brand
	if err := bm.db.Find(&brand).Error; err != nil {
		return nil, err
	}
	return brand, nil
}

func (bm *GormBrandModel) Insert(brand Brand) (Brand, error) {
	tx := bm.db.Begin()
	if err := tx.Create(&brand).Error; err != nil {
		tx.Rollback()
		return brand, err
	}
	tx.Commit()
	return brand, nil
}

func (bm *GormBrandModel) Edit(id int, brand Brand) (Brand, error) {
	tx := bm.db.Begin()
	if err := tx.Model(Brand{}).Where("id=?", id).Updates(Brand{
		Name: brand.Name, Logo: brand.Logo, Banner: brand.Banner,
	}).Error; err != nil {
		tx.Rollback()
		return Brand{}, err
	}
	if err := tx.First(&brand, id).Error; err != nil {
		tx.Rollback()
		return Brand{}, err
	}
	tx.Commit()
	return brand, nil
}

func (bm *GormBrandModel) Delete(id int) (Brand, error) {
	var brand Brand
	tx := bm.db.Begin()
	if err := tx.First(&brand, id).Error; err != nil {
		tx.Rollback()
		return Brand{}, err
	}
	if err := tx.Delete(&Brand{}, id).Error; err != nil {
		tx.Rollback()
		return Brand{}, err
	}
	tx.Commit()
	return brand, nil
}
