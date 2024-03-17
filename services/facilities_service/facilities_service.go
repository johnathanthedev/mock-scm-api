package facilities_service

import (
	models "scm-api/api/models"
	"scm-api/db"
	facility_dtos "scm-api/types/facilities/dtos"
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
