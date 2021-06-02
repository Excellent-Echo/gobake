package product

import (
	"go-bake/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Product, error)
	FindByID(ID string) (entity.Product, error)
	FindByName(name string) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Product, error)
	DeleteByID(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.Product, error) {

	var products []entity.Product

	if err := r.db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindByID(ID string) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("id=?", ID).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByName(name string) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("name = ?", name).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Create(product entity.Product) (entity.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Model(&product).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return product, err
	}

	if err := r.db.Where("id = ?", ID).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {

	if err := r.db.Where("id = ?", ID).Delete(&entity.Product{}).Error; err != nil {
		return "error", err
	}
	return "success delete", nil
}
