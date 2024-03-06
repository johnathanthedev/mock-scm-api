package operations_service

import (
	"errors"
	models "scm-api/api/models"
	operation_types "scm-api/types/operations/requests"

	"scm-api/db"

	"gorm.io/gorm"
)

func OperationExists(name string) bool {
	var existingOperation models.Operation
	result := db.GetDB().Where("name = ?", name).First(&existingOperation)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CreateOperation(operation operation_types.CreateOperationRequest) (*models.Operation, error) {
	newOperation := &models.Operation{
		Name:   operation.Name,
		Status: operation.Status,
	}

	if err := db.GetDB().Create(newOperation).Error; err != nil {
		return nil, err
	}

	return newOperation, nil
}
