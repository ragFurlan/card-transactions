package accounts

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
	"fmt"
)

type AccountUseCase struct {
	AccountRepository repository.Accounts
}

type Account interface {
	Save(account entity.Account) error
	Get(id int) (entity.Account, error)
}

func NewAccountsUseCase(accountRepository repository.Accounts) *AccountUseCase {
	return &AccountUseCase{
		AccountRepository: accountRepository,
	}
}

func (a AccountUseCase) Save(account entity.Account) error {
	accountByDocumentNumber, err := a.GetByDocumentNumber(account.DocumentNumber)
	if err != nil {
		return err
	}

	if accountByDocumentNumber.ID != "" {
		return fmt.Errorf("This account has already been created")
	}

	err = a.AccountRepository.Save(account)
	if err != nil {
		return err
	}

	return nil
}

func (a AccountUseCase) GetByID(id string) (entity.Account, error) {
	account, err := a.AccountRepository.GetByID(id)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (n AccountUseCase) GetByDocumentNumber(documentNumber string) (entity.Account, error) {
	account, err := n.AccountRepository.GetByDocumentNumber(documentNumber)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}