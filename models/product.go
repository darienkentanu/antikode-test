package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `gorm:"type:varchar(200);unique;not null" json:"name"`
	Picture string  `gorm:"type:longtext" json:"picture"`
	Price   float32 `gorm:"type:decimal(20,2);not null" json:"price"`
	BrandID uint    `gorm:"type:bigint;not null" json:"brand_id"`
}

type GormProductModel struct {
	db *gorm.DB
}

func NewProductModel(db *gorm.DB) *GormProductModel {
	return &GormProductModel{db: db}
}

type ProductModel interface {
	Insert(product Product) (Product, error)
	GetAll() ([]Product, error)
	Edit(id int, product Product) (Product, error)
	Delete(id int) (Product, error)
	GetProductById(id int) (Product, error)
}

func (pm *GormProductModel) Insert(product Product) (Product, error) {
	tx := pm.db.Begin()
	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return Product{}, err
	}
	tx.Commit()
	return product, nil
}

func (pm *GormProductModel) GetAll() ([]Product, error) {
	var products []Product
	tx := pm.db.Begin()
	if err := tx.Find(&products).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return products, nil
}

func (pm *GormProductModel) GetProductById(id int) (Product, error) {
	var product Product
	tx := pm.db.Begin()
	if err := tx.First(&product, id).Error; err != nil {
		tx.Rollback()
		return Product{}, err
	}
	tx.Commit()
	return product, nil
}

func (pm *GormProductModel) Edit(id int, product Product) (Product, error) {
	tx := pm.db.Begin()
	if err := tx.Model(&Product{}).Where(id).Updates(&product).Error; err != nil {
		tx.Rollback()
		return Product{}, err
	}
	tx.Commit()
	return product, nil
}

func (pm *GormProductModel) Delete(id int) (Product, error) {
	var product Product
	tx := pm.db.Begin()
	if err := tx.First(&product, id).Error; err != nil {
		tx.Rollback()
		return Product{}, err
	}
	if err := tx.Delete(&Product{}, id).Error; err != nil {
		tx.Rollback()
		return Product{}, err
	}
	tx.Commit()
	return product, nil
}
