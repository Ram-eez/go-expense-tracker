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
	words := strings.Split(message, " ")

	for _, word := range words {
		if strings.ToLower(word) != "debited" {
			break
		}
	}

	var amountFloat float64
	for _, word := range words {
		amount, err := strconv.ParseFloat(word, 64)
		if err == nil {
			amountFloat = amount
			break
		}
	}

	fmt.Println(amountFloat)

}
