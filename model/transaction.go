package model

type Transaction struct {
	Id         string `json:"id"`
	SenderId   string `json:"sender_id" db:"sender_id"`
	ReceiverId string `json:"receiver_id" db:"receiver_id"`
	Created_at string `db:"created_at" json:"created_at"`
}

// type TransactionTest struct {
// 	Id         string       `json:"id"`
// 	TotalPrice int          `json:"total" db:"total_price"`
// 	Created_at string       `db:"created_at" json:"created_at"`
// 	Updated_at sql.NullTime `db:"updated_at" json:"updated_at,omitempty"`
// }
