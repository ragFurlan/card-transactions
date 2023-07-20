package entity

import "time"

type Transactions struct {
	ID               int
	AccountID        int
	OperationTypesID int
	Account          float64
	EventDate        time.Time
}
