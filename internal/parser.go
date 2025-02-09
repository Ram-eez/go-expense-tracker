package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	id       int64
	amount   float64
	date     time.Time
	merchant string
}

func ParseMessage(message string) {
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
		return
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

	fmt.Println(merchant)

	fmt.Println(amountFloat)
}
