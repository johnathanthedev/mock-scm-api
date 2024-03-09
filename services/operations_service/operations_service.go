package operations_service

import (
	"errors"
	models "scm-api/api/models"
	operation_types "scm-api/types/operations/requests"

	"scm-api/db"

	"github.com/google/uuid"
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

func AddUserToOperation(operationID uuid.UUID, userID uuid.UUID) error {
	database := db.GetDB()

	var operation models.Operation

	if err := database.First(&operation, "id = ?", operationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("operation not found")
		}
		return err
	}

	if err := database.Model(&operation).Association("Users").Append(&models.User{ID: userID}); err != nil {
		return err
	}

	return nil
}

func IsUserInOperation(operationID uuid.UUID, userID uuid.UUID) (bool, error) {
	database := db.GetDB()

	exists := false

	query := `SELECT EXISTS(SELECT 1 FROM operation_users WHERE operation_id = ? AND user_id = ?)`

	err := database.Raw(query, operationID, userID).Row().Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func ListOperations() ([]models.Operation, error) {
	var operations []models.Operation
	result := db.GetDB().Find(&operations)

	if result.Error != nil {
		return nil, result.Error
	}

	return operations, nil
}
