package product

import (
	"go-bake/entity"
	"time"
)

type ProductFormat struct {
	ID          uint32 `json:"product_id"`
	ProductName string `json:"product_name"`
	Stock       string `json:"stock"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatProduct(product entity.Product) ProductFormat {

	var formatProduct = ProductFormat{
		ID:          product.ID,
		ProductName: product.ProductName,
		Stock:       product.Stock,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	}
	return formatProduct
}

func FormatDeleteProduct(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
