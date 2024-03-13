package vehicles_service

import (
	"errors"
	models "scm-api/api/models"
	vehicle_types "scm-api/types/vehicles/requests"

	"scm-api/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func VehicleExists(name string) bool {
	var existingVehicle models.Vehicle
	result := db.GetDB().Where("name = ?", name).First(&existingVehicle)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CreateVehicle(vehicle vehicle_types.CreateVehicleRequest) (*models.Vehicle, error) {
	newVehicle := &models.Vehicle{
		OperationID:    vehicle.OperationID,
		VehicleType:    vehicle.VehicleType,
		Name:           vehicle.Name,
		CarryVolume:    vehicle.CarryVolume,
		MaxWeight:      vehicle.MaxWeight,
		LastLocation:   vehicle.LastLocation,
		Status:         vehicle.Status,
		PreferredSpeed: vehicle.PreferredSpeed,
		CrewCapacity:   vehicle.CrewCapacity,
		Attributes:     vehicle.Attributes,
	}

	if err := db.GetDB().Create(newVehicle).Error; err != nil {
		return nil, err
	}

	return newVehicle, nil
}

func GetAllVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	result := db.GetDB().Find(&vehicles)

	if result.Error != nil {
		return nil, result.Error
	}

	return vehicles, nil
}

func GetVehicleById(vehicleID uuid.UUID) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	result := db.GetDB().First(&vehicle, "id = ?", vehicleID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("vehicle not found")
		}
		return nil, result.Error
	}

	return &vehicle, nil
}
