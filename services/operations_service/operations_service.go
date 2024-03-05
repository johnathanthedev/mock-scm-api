package operations_service

import (
	"errors"
	dtos "scm-api/api/dtos/requests"
	models "scm-api/api/models"

	"scm-api/db"

	"gorm.io/gorm"
)

func OperationExists(name string) bool {
	var existingOperation models.Operation
	result := db.GetDB().Where("name = ?", name).First(&existingOperation)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CreateOperation(operation dtos.CreateOperationRequest) (*models.Operation, error) {
	newOperation := &models.Operation{
		Name:   operation.Name,
		Status: operation.Status,
	}

	if err := db.GetDB().Create(newOperation).Error; err != nil {
		return nil, err
	}

	return newOperation, nil
}
