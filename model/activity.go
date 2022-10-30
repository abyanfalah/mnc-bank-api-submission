package model

import "time"

type Activity struct {
	Id         string    `json:"id"`
	CustomerId string    `json:"customer_id"`
	Activity   string    `json:"activity"`
	Time       time.Time `json:"time"`
}
