package product

import (
	"errors"
	"fmt"
	"go-bake/entity"
	"go-bake/helper"
	"time"
)

type Service interface {
	GetAllProduct() ([]ProductFormat, error)
	GetProductByID(ID string) (ProductFormat, error)
	SaveNewProduct(product entity.ProductInput) (ProductFormat, error)
	UpdateProductByID(ID string, dataInput entity.UpdateProductInput) (ProductFormat, error)
	DeleteProductByID(ID string) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetAllProduct() ([]ProductFormat, error) {
	products, err := s.repository.FindAll()

	var formatProducts []ProductFormat

	for _, product := range products {
		formatProduct := FormatProduct(product)
		formatProducts = append(formatProducts, formatProduct)
	}

	if err != nil {
		return formatProducts, err
	}
	return formatProducts, nil
}

func (s *service) GetProductByID(ID string) (ProductFormat, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return ProductFormat{}, err
	}

	product, err := s.repository.FindByID(ID)

	if err != nil {
		return ProductFormat{}, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %d not found", product.ID)

		return ProductFormat{}, errors.New(newError)
	}

	formatProduct := FormatProduct(product)
	return formatProduct, nil
}

func (s *service) SaveNewProduct(product entity.ProductInput) (ProductFormat, error) {

	var newProduct = entity.Product{

		ProductName: product.ProductName,
		Stock:       product.Stock,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	createProduct, err := s.repository.Create(newProduct)
	formatProduct := FormatProduct(createProduct)

	if err != nil {
		return formatProduct, err
	}

	return formatProduct, nil
}

func (s *service) UpdateProductByID(ID string, dataInput entity.UpdateProductInput) (ProductFormat, error) {

	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return ProductFormat{}, err
	}

	product, err := s.repository.FindByID(ID)

	if err != nil {
		return ProductFormat{}, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %d not found", product.ID)
		return ProductFormat{}, errors.New(newError)
	}

	if dataInput.ProductName != "" || len(dataInput.ProductName) != 0 {
		dataUpdate["product_name"] = dataInput.ProductName
	}

	if dataInput.Stock != "" || len(dataInput.Stock) != 0 {
		dataUpdate["stock"] = dataInput.Stock
	}

	if dataInput.Price != "" || len(dataInput.Price) != 0 {
		dataUpdate["price"] = dataInput.Price
	}

	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}

	if dataInput.Image != "" || len(dataInput.Image) != 0 {
		dataUpdate["image"] = dataInput.Image
	}

	dataUpdate["updated_at"] = time.Now()

	productUpdate, err := s.repository.UpdateByID(ID, dataUpdate)
	if err != nil {
		return ProductFormat{}, err
	}

	formatProduct := FormatProduct(productUpdate)

	return formatProduct, nil
}

func (s *service) DeleteProductByID(ID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return nil, err
	}

	product, err := s.repository.FindByID(ID)

	if err != nil {
		return nil, err
	}

	if product.ID == 0 {
		newError := fmt.Sprintf("product id %d not found", product.ID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(ID)

	if err != nil {
		return nil, err
	}

	if status == "error" {
		return nil, errors.New("error delete")
	}

	msg := fmt.Sprintf("succes delete product ID : %s", ID)

	formatDelete := FormatDeleteProduct(msg)

	return formatDelete, nil

}
