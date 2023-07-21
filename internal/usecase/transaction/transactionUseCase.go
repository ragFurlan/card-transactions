package transaction

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
	accountsUseCase "card-transactions/internal/usecase/accounts"
	"fmt"
)

type TransactionUseCase struct {
	TransactionRepository   repository.Transaction
	OperationTypeRepository repository.OperationType
	AccountUseCase          accountsUseCase.AccountUseCase
}

type Transaction interface {
	Save(account entity.Transaction) error
}

func NewTransactionUseCase(transactionRepository repository.Transaction,
	operationTypeRepository repository.OperationType,
	accountUseCase accountsUseCase.AccountUseCase) *TransactionUseCase {

	return &TransactionUseCase{
		TransactionRepository:   transactionRepository,
		OperationTypeRepository: operationTypeRepository,
		AccountUseCase:          accountUseCase,
	}
}

func (t TransactionUseCase) Save(transaction entity.Transaction) error {
	account, err := t.AccountUseCase.GetByID(transaction.AccountID)
	if err != nil {
		return err
	}

	if account.ID == "" {
		return fmt.Errorf("The given account does not exist")
	}

	operationType, err := t.OperationTypeRepository.GetByID(transaction.OperationTypesID)
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
