package repository

import (
	"bufio"
	"card-transactions/internal/entity"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type TransactionRepository struct {
	transactionFilePath string
}

type Transaction interface {
	Save(transactions entity.Transaction) error
	GetTransactionByAccountID(accountID string) ([]entity.Transaction, error)
}

func NewTransactionRepository(transactionFilePath string) Transaction {
	return &TransactionRepository{
		transactionFilePath: transactionFilePath,
	}
}

func (r TransactionRepository) Save(transactions entity.Transaction) error {
	file, err := os.OpenFile(r.transactionFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open transaction file: %v", err)
	}
	defer file.Close()

	transactionEntry := fmt.Sprintf("ID: %v|AccountID: %s|OperationTypesID: %v|Amount: %v|EventDate: %v\n",
		transactions.ID,
		transactions.AccountID,
		transactions.OperationTypesID,
		transactions.Amount,
		transactions.EventDate)

	if _, err := file.WriteString(transactionEntry); err != nil {
		return fmt.Errorf("Failed to write account entry: %v", err)
	}
	return nil
}

func (r TransactionRepository) GetTransactionByAccountID(accountID string) ([]entity.Transaction, error) {
	file, err := os.Open(r.transactionFilePath)
	if err != nil {
		err := fmt.Errorf("Failed to open transaction file: %v", err)
		return []entity.Transaction{}, err
	}
	defer file.Close()

	transactions := []entity.Transaction{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		transactionEntry := scanner.Text()
		t, err := parseTransactionEntry(transactionEntry)
		if err != nil {
			fmt.Printf("Failed to parse account entry: %v", err)
			continue
		}

		if t.AccountID == accountID {
			transactions = append(transactions, t)
		}
	}

	return transactions, nil
}

func parseTransactionEntry(transactionEntry string) (entity.Transaction, error) {
	var transaction entity.Transaction

	lines := strings.Split(transactionEntry, "|")
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			return transaction, fmt.Errorf("invalid transaction entry format: %s", transactionEntry)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "EventDate":
			eventDate, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return transaction, fmt.Errorf("failed to parse EventDate in transaction entry: %s", transactionEntry)
			}
			transaction.EventDate = eventDate
		case "OperationTypesID":
			operationTypesID, _ := strconv.Atoi(value)
			transaction.OperationTypesID = operationTypesID
		case "ID":
			transaction.ID = value
		case "AccountID":
			transaction.AccountID = value
		case "Amount":
			amount, _ := strconv.ParseFloat(value, 64)
			transaction.Amount = amount

		}
	}

	return transaction, nil
}
