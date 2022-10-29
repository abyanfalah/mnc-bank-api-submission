package model

import "time"

type Activity struct {
	Id         string
	CustomerId string
	Activity   string
	Time       time.Time
}
