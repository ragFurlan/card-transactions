package transaction

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
	accountsUseCase "card-transactions/internal/usecase/accounts"
	operationTypeUseCase "card-transactions/internal/usecase/operationType"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type TransactionUseCase struct {
	TransactionRepository repository.Transaction
	OperationTypeUseCase  operationTypeUseCase.OperationTypeUseCase
	AccountUseCase        accountsUseCase.AccountUseCase
}

type Transaction interface {
	Save(account entity.Transaction) error
}

func NewTransactionUseCase(transactionRepository repository.Transaction,
	operationTypeUseCase operationTypeUseCase.OperationTypeUseCase,
	accountUseCase accountsUseCase.AccountUseCase) *TransactionUseCase {

	return &TransactionUseCase{
		TransactionRepository: transactionRepository,
		OperationTypeUseCase:  operationTypeUseCase,
		AccountUseCase:        accountUseCase,
	}
}

func (t TransactionUseCase) Save(transaction entity.Transaction) error {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return err
	}

	account, err := t.AccountUseCase.GetByID(transaction.AccountID)
	if err != nil {
		return err
	}

	if account.ID == "" {
		return fmt.Errorf("The given account does not exist")
	}

	operationType, err := t.OperationTypeUseCase.Get(transaction.OperationTypesID)
	if err != nil {
		return err
	}

	if operationType.ID == 0 {
		return fmt.Errorf("The given operation type does not exist")
	}

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return err
	}

	return nil
}
