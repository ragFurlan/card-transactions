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
	AccountUsecase *accounts.AccountUsecase
}

func NewAccountHandler(accountUsecase *accounts.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		AccountUsecase: accountUsecase,
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

	account, err := h.AccountUsecase.GetByID(id)
	if err != nil {
	
		http.Error(w, "Failed to get account", http.StatusInternalServerError)
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

	err = h.AccountUsecase.Save(account)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		AccountID string `json:"account_id"`
		Message   string `json:"message"`
	}{
		AccountID: account.ID,
		Message:   "Save account with success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
