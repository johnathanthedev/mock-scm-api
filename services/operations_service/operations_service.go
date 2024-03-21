package operations_service

import (
	"errors"
	models "scm-api/api/models"
	operation_dtos "scm-api/types/operations/dtos"
	operation_types "scm-api/types/operations/requests"
	"sort"

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

func ListUserJoinedOperations(userID uuid.UUID) ([]models.Operation, error) {
	var operations []models.Operation
	database := db.GetDB()

	err := database.Table("operations").
		Joins("JOIN operation_users on operations.id = operation_users.operation_id").
		Joins("JOIN users on users.id = operation_users.user_id").
		Where("users.id = ?", userID).
		Scan(&operations).Error

	if err != nil {
		return nil, err
	}

	if len(operations) == 0 {
		return []models.Operation{}, nil
	}

	return operations, nil
}

func ListOperationsWithJoinStatus(userID uuid.UUID) ([]operation_dtos.OperationDto, error) {
	var operations []models.Operation // Use the correct model that maps to your DB table
	database := db.GetDB()

	// First, get all operations from the 'operations' table.
	if err := database.Find(&operations).Error; err != nil {
		return nil, err
	}

	// Create a map to hold join status for operations.
	operationsJoined := make(map[uuid.UUID]bool)

	// Fetch all operations the user has joined and mark them in the map.
	var userOperations []models.Operation
	if err := database.Table("operations").
		Joins("JOIN operation_users on operations.id = operation_users.operation_id").
		Where("operation_users.user_id = ?", userID).
		Select("operations.*"). // Make sure to select fields from the 'operations' table
		Scan(&userOperations).Error; err != nil {
		return nil, err
	}

	for _, op := range userOperations {
		operationsJoined[op.ID] = true
	}

	// Map operations to OperationDto, updating the Joined field based on the operationsJoined map.
	operationsDtos := make([]operation_dtos.OperationDto, len(operations))
	for i, op := range operations {
		joined := operationsJoined[op.ID]
		operationsDtos[i] = operation_dtos.OperationDto{
			ID:     op.ID,
			Name:   op.Name,
			Status: op.Status,
			Joined: joined,
		}
	}

	// Sort the operationsDtos slice by the Name field.
	sort.Slice(operationsDtos, func(i, j int) bool {
		return operationsDtos[i].Name < operationsDtos[j].Name
	})

	return operationsDtos, nil
}

func RemoveUserFromOperation(operation models.Operation, userID uuid.UUID) error {
	database := db.GetDB()

	user := models.User{ID: userID}

	if err := database.Model(&operation).Association("Users").Delete(&user); err != nil {
		return err
	}

	return nil
}

func GetOperationByID(operationID uuid.UUID) (*models.Operation, error) {
	var operation models.Operation
	result := db.GetDB().First(&operation, "id = ?", operationID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("operation not found")
		}
		return nil, result.Error
	}

	return &operation, nil
}
