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
	// dsn taking postgres as DB
	// dsn := "host=localhost user=postgres password= dbname=golang_demo port=5432 sslmode=disable"
	dbURL := "postgres://postgres:password@localhost:5432/golang_demo?sslmode=disable"
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		panic("DB connection failed!!!!!")
	}

	// auto migration
	// if err = database.AutoMigrate(&domain.Tenant{}); err != nil {
	// 	log.Println(err)
	// }

	DB = database
	fmt.Println("DB connection successful!!!!!")
}
