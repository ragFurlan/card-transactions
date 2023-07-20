package repository

import (
	"bufio"
	"card-transactions/internal/entity"
	"fmt"
	"os"
	"strings"
)

type AccountsRepository struct {
	accountFilePath string
}

type Accounts interface {
	Save(account entity.Account) error
	GetByID(id string) (entity.Account, error)
	GetByDocumentNumber(documentNumber string) (entity.Account, error)
}

func NewAccountsRepository(accountFilePath string) Accounts {
	return &AccountsRepository{
		accountFilePath: accountFilePath,
	}
}

func (r AccountsRepository) Save(account entity.Account) error {
	file, err := os.OpenFile(r.accountFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open account file: %v", err)
	}
	defer file.Close()

	accountEntry := fmt.Sprintf("ID: %v|DocumentNumber: %s\n", account.ID, account.DocumentNumber)

	if _, err := file.WriteString(accountEntry); err != nil {
		return fmt.Errorf("Failed to write account entry: %v", err)
	}
	return nil

}

func (r AccountsRepository) GetByID(id string) (entity.Account, error) {
	file, err := os.Open(r.accountFilePath)
	if err != nil {
		err := fmt.Errorf("Failed to open account file: %v", err)
		return entity.Account{}, err
	}
	defer file.Close()

	account := entity.Account{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		accountEntry := scanner.Text()
		a, err := parseAccountEntry(accountEntry)
		if err != nil {
			fmt.Printf("Failed to parse account entry: %v", err)
			continue
		}

		if a.ID == id {
			account = a
		}
	}

	return account, nil
}

func (r AccountsRepository) GetByDocumentNumber(documentNumber string) (entity.Account, error) {
	file, err := os.Open(r.accountFilePath)
	if err != nil {
		err := fmt.Errorf("Failed to open account file: %v", err)
		return entity.Account{}, err
	}
	defer file.Close()

	account := entity.Account{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		accountEntry := scanner.Text()
		a, err := parseAccountEntry(accountEntry)
		if err != nil {
			fmt.Printf("Failed to parse account entry: %v", err)
			continue
		}

		if a.DocumentNumber == documentNumber {
			account = a
		}
	}

	return account, nil
}

func parseAccountEntry(accountEntry string) (entity.Account, error) {
	var account entity.Account

	lines := strings.Split(accountEntry, "|")
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			return account, fmt.Errorf("invalid account entry format: %s", accountEntry)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "DocumentNumber":
			account.DocumentNumber = value
		case "ID":
			account.ID = value
		}
	}

	return account, nil
}
