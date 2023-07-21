package entity

type Account struct {
	ID             string `json:"account_id" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required"`
}
