package operationType

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
)

type OperationTypeUsecase struct {
	OperationTypeRepository repository.OperationType
}

type OperationType interface {
	Get(id int) (entity.OperationType, error)
}

func NewOperationTypeUsecase(operationTypeRepository repository.OperationType) *OperationTypeUsecase {
	return &OperationTypeUsecase{
		OperationTypeRepository: operationTypeRepository,
	}
}

func (r OperationTypeUsecase) Get(id int) (entity.OperationType, error) {
	account, err := r.OperationTypeRepository.GetByID(id)
	if err != nil {
		return entity.OperationType{}, err
	}

	return account, nil
}
