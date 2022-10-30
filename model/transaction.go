package model

import "time"

type Transaction struct {
	Id         string    `json:"id"`
	SenderId   string    `json:"sender_id" db:"sender_id"`
	ReceiverId string    `json:"receiver_id" db:"receiver_id"`
	Amount     int       `json:"amount" db:"amount"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}
