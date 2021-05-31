package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint32         `gorm:"primarykey" json: "id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Username  string         `json:"username"`
	Phone     uint           `json:"phone"`
	Address   string         `json:"adress"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
