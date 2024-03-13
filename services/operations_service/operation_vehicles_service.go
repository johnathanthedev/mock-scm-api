package operations_service

import (
	"errors"
	models "scm-api/api/models"
	"scm-api/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddVehicleToOperation(operation *models.Operation, vehicle *models.Vehicle) error {
	database := db.GetDB()

	if err := database.Model(operation).Association("Vehicles").Append(vehicle); err != nil {
		return err
	}

	return nil
}

func IsVehicleInOperation(operationID uuid.UUID, vehicleID uuid.UUID) bool {
	database := db.GetDB()

	var vehicle models.Vehicle

	result := database.Where("id = ? AND operation_id = ?", vehicleID, operationID).First(&vehicle)

	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}
