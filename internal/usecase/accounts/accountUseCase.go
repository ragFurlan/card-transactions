package accounts

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type AccountUsecase struct {
	AccountRepository repository.Accounts
}

type Account interface {
	Save(account entity.Account) error
	Get(id int) (entity.Account, error)
}

func NewAccountsUsecase(accountRepository repository.Accounts) *AccountUsecase {
	return &AccountUsecase{
		AccountRepository: accountRepository,
	}
}

func (a AccountUsecase) Save(account entity.Account) error {
	validate := validator.New()
	err := validate.Struct(account)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

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

func (a AccountUsecase) GetByID(id string) (entity.Account, error) {
	account, err := a.AccountRepository.GetByID(id)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (n AccountUsecase) GetByDocumentNumber(documentNumber string) (entity.Account, error) {
	account, err := n.AccountRepository.GetByDocumentNumber(documentNumber)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}
