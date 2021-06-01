package entity

type Product struct {
	ID          uint32 `json:"product_id"`
	ProductName string `json:"product_name"`
	Stock       int    `json:"stock"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
