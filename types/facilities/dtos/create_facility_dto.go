package facility_dtos

import (
	models "scm-api/api/models/geo"

	"github.com/google/uuid"
)

type CreateFacilityDto struct {
	Name               string          `json:"name" validate:"required"`
	Type               string          `json:"type" validate:"required"`
	MaxStorageCapacity int             `json:"max_storage_capacity" validate:"required"`
	DailyOperatingCost int             `json:"daily_operating_cost" validate:"required"`
	DailyRentCost      int             `json:"daily_rent_cost" validate:"required"`
	DailyCarbonOutput  int             `json:"daily_carbon_output" validate:"required"`
	Location           models.GeoPoint `json:"location" validate:"required"`
	OperationID        uuid.UUID       `json:"operation_id" validate:"required"`
}
