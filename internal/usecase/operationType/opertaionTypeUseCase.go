package operationType

import (
	"card-transactions/internal/entity"
	repository "card-transactions/internal/platform/repositories"
)

type OperationTypeUseCase struct {
	OperationTypeRepository repository.OperationType
}

type OperationType interface {
	Get(id int) (entity.OperationType, error)
}

func NewOperationTypeUseCase(operationTypeRepository repository.OperationType) *OperationTypeUseCase {
	return &OperationTypeUseCase{
		OperationTypeRepository: operationTypeRepository,
	}
}

func (r OperationTypeUseCase) Get(id int) (entity.OperationType, error) {
	account, err := r.OperationTypeRepository.GetByID(id)
	if err != nil {
		return entity.OperationType{}, err
	}

	return account, nil
}
