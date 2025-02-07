package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env")
	}
}

func Connect() {
	dsn := os.Getenv("DB")
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	db = d
}
