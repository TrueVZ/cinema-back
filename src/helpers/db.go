package helpers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func CreateDatabaseInstance() *gorm.DB {
	dsn := "host=database user=admin password=hackme dbname=cinema port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[ ERROR ] Unable to connect with mysql!\n", err)
	}

	fmt.Println("[ OK ] Connected to the DB!")

	return db
}
