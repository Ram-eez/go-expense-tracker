package internal

import "time"

type Message struct {
	Id       int64
	Amount   float64
	Date     time.Time
	Merchant string
}
