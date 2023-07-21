package handler

import (
	"card-transactions/internal/entity"
	"card-transactions/internal/usecase/transaction"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TransactionHandler struct {
	TransactionUseCase *transaction.TransactionUseCase
}

func NewTransactionHandler(transactionUseCase *transaction.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		TransactionUseCase: transactionUseCase,
	}
}

func (h *TransactionHandler) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		AccountID        string  `json:"account_id"`
		OperationTypesID int     `json:"operation_type_id"`
		Amount           float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	uuid := uuid.New()
	transaction := entity.Transaction{
		ID:               uuid.String(),
		AccountID:        requestBody.AccountID,
		OperationTypesID: requestBody.OperationTypesID,
		Amount:           requestBody.Amount,
		EventDate:        time.Now(),
	}

	err = h.TransactionUseCase.Save(transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Save transaction with success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
