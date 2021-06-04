package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint32         `json:"product_id"`
	ProductName string         `json:"product_name"`
	Stock       int            `json:"stock"`
	Price       float64        `json:"price"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"update_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type ProductInput struct {
	ProductName string  `json:"product_name" form:"product_name"`
	Stock       int     `json:"stock" form:"stock"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
}

type UpdateProductInput struct {
	ProductName string  `json:"product_name" form:"product_name"`
	Stock       int     `json:"stock" form:"stock"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
}
