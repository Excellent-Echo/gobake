package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint32         `json:"product_id"`
	ProductName string         `json:"product_name"`
	Stock       string         `json:"stock"`
	Price       string         `json:"price"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"update_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type ProductInput struct {
	ProductName string `json:"product_name"`
	Stock       string `json:"stock"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type UpdateProductInput struct {
	ProductName string `json:"product_name"`
	Stock       string `json:"stock"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
