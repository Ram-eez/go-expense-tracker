package internal

import (
	"expensetracker/config"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type Message struct {
	Id       int64     `json:"id"`
	Amount   float64   `json:"amount"`
	Date     time.Time `json:"date"`
	Merchant string    `json:"merchant"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	if db == nil {
		log.Fatal("Database conn Failed")
	} else {
		log.Println("Database conn Successfull")
	}
}

func (m *Message) CreateTransaction() error {
	if err := db.Create(m).Error; err != nil {
		fmt.Println("failed to create the transaction")
		return err
	}
	return nil
}
