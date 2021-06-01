package config

import (
	"fmt"
	"go-bake/entity"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//creating a new connection to our database
func SetupConnection() *gorm.DB {
	err := godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.Transaction{})
	db.AutoMigrate(&entity.Cart{})
	return db
}
