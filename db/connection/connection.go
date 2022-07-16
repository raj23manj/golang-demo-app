package connection

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Database() {
	dsn := "host=localhost user=postgres password= dbname=golang_demo port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		panic("DB connection failed!!!!!")
	}
	DB = database
	fmt.Println("DB connection successful!!!!!")
	return
}
