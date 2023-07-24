package transaction

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
	accountsUsecase "card-transactions/internal/usecase/accounts"
	operationTypeUsecase "card-transactions/internal/usecase/operationType"
	"fmt"

	"gopkg.in/go-playground/validator.v9"

)

type TransactionUsecase struct {
	TransactionRepository repository.Transaction
	OperationTypeUsecase  operationTypeUsecase.OperationTypeUsecase
	AccountUsecase        accountsUsecase.AccountUsecase
}

type Transaction interface {
	Save(account entity.Transaction) error
}

func NewTransactionUsecase(transactionRepository repository.Transaction,
	operationTypeUsecase operationTypeUsecase.OperationTypeUsecase,
	accountUsecase accountsUsecase.AccountUsecase) *TransactionUsecase {

	return &TransactionUsecase{
		TransactionRepository: transactionRepository,
		OperationTypeUsecase:  operationTypeUsecase,
		AccountUsecase:        accountUsecase,
	}
}

func (t TransactionUsecase) Save(transaction entity.Transaction) error {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return err
	}

	account, err := t.AccountUsecase.GetByID(transaction.AccountID)
	if err != nil {
		return err
	}

	if account.ID == "" {
		return fmt.Errorf("The given account does not exist")
	}

	operationType, err := t.OperationTypeUsecase.Get(transaction.OperationTypesID)
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
