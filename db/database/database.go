package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password= dbname=golang_demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		panic("DB connection failed!!!!!")
	}

	fmt.Println("DB connection successful!!!!!")
	return db, err
}
