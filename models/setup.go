package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=venus123"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := database.AutoMigrate(&Book{}); err != nil {
		fmt.Println(err)
	}

	DB = database

	fmt.Println("Conexão feita com sucesso.")
}
