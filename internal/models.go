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

type MessageList struct {
	Messages []Message `gorm:"-"`
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

func (m *Message) CreateTransactionDB() error {
	if err := db.Create(m).Error; err != nil {
		fmt.Println("failed to create the transaction")
		return err
	}
	return nil
}

func (ml *MessageList) GetAllTransactionsFromDB() error {
	if err := db.Find(&ml.Messages).Error; err != nil {
		log.Println("Failed to fetch all transactions")
		return err
	}
	return nil
}

func (m *Message) GetTransactionDB() error {
	if err := db.First(m, m.Id).Error; err != nil {
		log.Println("Unable to fetch the trasnaction")
		return err
	}
	return nil
}
