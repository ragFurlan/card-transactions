package repository

import (
	"bufio"
	"card-transactions/internal/entity"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OperationTypeRepository struct {
	operationTypeFilePath string
}

type OperationType interface {
	GetByID(id int) (entity.OperationType, error)
}

func NewOperationTypeRepository(operationTypeFilePath string) OperationType {
	return &OperationTypeRepository{
		operationTypeFilePath: operationTypeFilePath,
	}
}

func (r OperationTypeRepository) GetByID(id int) (entity.OperationType, error) {
	file, err := os.Open(r.operationTypeFilePath)
	if err != nil {
		err := fmt.Errorf("Failed to open operation type file: %v", err)
		return entity.OperationType{}, err
	}
	defer file.Close()

	operationType := entity.OperationType{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		operationTypeEntry := scanner.Text()
		a, err := parseOperationTypeEntry(operationTypeEntry)
		if err != nil {
			fmt.Printf("Failed to parse operation type entry: %v", err)
			continue
		}

		if a.ID == id {
			operationType = a
		}
	}

	return operationType, nil
}

func parseOperationTypeEntry(operationTypeEntry string) (entity.OperationType, error) {
	var operationType entity.OperationType

	lines := strings.Split(operationTypeEntry, "|")
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			return operationType, fmt.Errorf("invalid operation type entry format: %s", operationTypeEntry)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Description":
			operationType.Description = value
		case "ID":
			id, _ := strconv.Atoi(value)
			operationType.ID = id
		}
	}

	return operationType, nil
}
