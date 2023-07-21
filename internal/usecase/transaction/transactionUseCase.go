package transaction

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
)

type TransactionUseCase struct {
	TransactionRepository repository.Transaction
}

type Transaction interface {
	Save(account entity.Transaction) error
}

func NewTransactionUseCase(transactionRepository repository.Transaction) *TransactionUseCase {
	return &TransactionUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (t TransactionUseCase) Save(transaction entity.Transaction) error {
	err := t.TransactionRepository.Save(transaction)
	if err != nil {
		return err
	}

	return nil
}
