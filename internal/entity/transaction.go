package entity

import "time"

type Transaction struct {
	ID               string    `json:"id"`
	AccountID        string    `json:"account_id"`
	OperationTypesID int       `json:"operation_type_id"`
	Amount           float64   `json:"amount"`
	EventDate        time.Time `json:"event_date"`
}
