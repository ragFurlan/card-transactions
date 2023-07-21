package handler

import (
	"card-transactions/internal/entity"
	"card-transactions/internal/usecase/accounts"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	AccountUseCase *accounts.AccountUseCase
}

func NewAccountHandler(accountUseCase *accounts.AccountUseCase) *AccountHandler {
	return &AccountHandler{
		AccountUseCase: accountUseCase,
	}
}

func (h *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["accountId"]
	if !ok {
		http.Error(w, "Id is missing in parameters", http.StatusMethodNotAllowed)
	}

	account, err := h.AccountUseCase.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to get logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		DocumentNumber string `json:"document_number"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	uuid := uuid.New()
	account := entity.Account{
		ID:             uuid.String(),
		DocumentNumber: requestBody.DocumentNumber,
	}

	err = h.AccountUseCase.Save(account)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Save account with success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
