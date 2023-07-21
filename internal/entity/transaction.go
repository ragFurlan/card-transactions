package entity

import "time"

type Transaction struct {
	ID               string    `json:"id" validate:"required"`
	AccountID        string    `json:"account_id" validate:"required"`
	OperationTypesID int       `json:"operation_type_id" validate:"required"`
	Amount           float64   `json:"amount" validate:"required"`
	EventDate        time.Time `json:"event_date" validate:"required"`
}
