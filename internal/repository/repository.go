package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	connectURL := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "user", "password", "test_db")
	d, err := gorm.Open("postgres", connectURL)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
