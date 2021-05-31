package entity

type Transaction struct {
	ID uint32 `gorm:"primaryKey" json:"transaction_id"`
	//Users `gorm:"foreignKey:"`
}
