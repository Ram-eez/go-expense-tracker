package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Id       int64
	Amount   float64
	Date     time.Time
	Merchant string
}

func ParseMessage(message string) *Message {
	fmt.Println("starting")
	words := strings.Split(message, " ")

	found := false
	for _, word := range words {
		fmt.Println(word)
		if strings.ToLower(word) == "debited" {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("No debited found")
		return nil
	}

	var amountFloat float64
	for _, word := range words {
		if value, err := strconv.ParseFloat(word, 64); err == nil {
			amountFloat = value
			break
		}
	}

	var merchant string
	for i, word := range words {
		if word == "to" {
			merchant = words[i+1]
		}
	}

	// var time string

	// for i, word := range words {
	// 	if word == "date" {
	// 		time = words[i+1]
	// 	}
	// }

	currentTime := time.Now()
	// currentTime.Format("2006-01-02 15:04:05")

	newMessage := &Message{
		Amount:   amountFloat,
		Date:     currentTime,
		Merchant: merchant,
	}

	return newMessage
}
