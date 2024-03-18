package facilities_service

import (
	"errors"
	models "scm-api/api/models"
	"scm-api/db"
	facility_dtos "scm-api/types/facilities/dtos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFacility(facilityDto facility_dtos.CreateFacilityDto) (*models.Facility, error) {
	newFacility := &models.Facility{
		Name:               facilityDto.Name,
		Type:               facilityDto.Type,
		MaxStorageCapacity: facilityDto.MaxStorageCapacity,
		DailyOperatingCost: facilityDto.DailyOperatingCost,
		DailyRentCost:      facilityDto.DailyRentCost,
		DailyCarbonOutput:  facilityDto.DailyCarbonOutput,
		Location:           facilityDto.Location,
		OperationID:        facilityDto.OperationID,
	}

	if err := db.GetDB().Create(newFacility).Error; err != nil {
		return nil, err
	}

	return newFacility, nil
}

func GetFacilityByID(facilityID uuid.UUID) (*models.Facility, error) {
	var facility models.Facility
	result := db.GetDB().First(&facility, "id = ?", facilityID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("facility not found")
		}
		return nil, result.Error
	}

	return &facility, nil
}
