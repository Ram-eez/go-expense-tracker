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
	Messages *[]Message `gorm:"-"`
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

func (ml *MessageList) GetAllTransactionsFromDB() error {
	if err := db.Find(&ml.Messages).Error; err != nil {
		log.Println("Failed to fetch all transactions")
		return err
	}
	return nil
}

func (ml *MessageList) SortByAmountAscDB() error {
	if err := db.Order("amount ASC").Find(&ml.Messages).Error; err != nil {
		log.Printf("Could not fetch the transactions : %d", err)
		return err
	}
	return nil
}

func (ml *MessageList) SortByAmountDescDB() error {
	if err := db.Order("amount DESC").Find(&ml.Messages).Error; err != nil {
		log.Printf("Could not fetch the transactions : %d", err)
		return err
	}
	return nil
}

func (ml *MessageList) SortByDateAscDB() error {
	if err := db.Order("date ASC").Find(&ml.Messages).Error; err != nil {
		log.Printf("Could not fetch the trasnactions : %d", err)
		return err
	}
	return nil
}

func (ml *MessageList) SortByDateDescDB() error {
	if err := db.Order("date DESC").Find(&ml.Messages).Error; err != nil {
		log.Printf("Could not fetch the trasnactions : %d", err)
		return err
	}
	return nil
}

func (ml *MessageList) TotalExpensesDB() (float64, error) {
	if err := db.Find(&ml.Messages).Error; err != nil {
		log.Printf("Could not fetch the trasnactions : %d", err)
		return -1, err
	}

	var TotalAmount float64
	for _, message := range *ml.Messages {
		TotalAmount += message.Amount
	}
	return TotalAmount, nil
}

func (m *Message) GetTransactionDB() error {
	if err := db.First(m, m.Id).Error; err != nil {
		log.Println("Unable to fetch the transaction")
		return err
	}
	return nil
}

func (m *Message) CreateTransactionDB() error {
	if err := db.Create(m).Error; err != nil {
		fmt.Println("failed to create the transaction")
		return err
	}
	return nil
}

func (m *Message) DeleteTransactionDB() error {
	if err := db.Where("id = ?", m.Id).Delete(m).Error; err != nil {
		log.Println("Could not delete the transaction")
		return err
	}
	return nil
}

func (m *Message) UpdateTransactionDB() error {
	if err := db.Model(m).Where("id = ?", m.Id).Updates(m).Error; err != nil {
		log.Printf("Could not update the trasnaction : %s", err)
		return err
	}
	return nil
}
