package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init opening a database and serve reference to `Database` struct
func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_HOST")), &gorm.Config{})
	if err != nil {
		log.Fatal("db err: (Init) ", err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
