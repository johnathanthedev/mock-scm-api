package route_dtos

import (
	"github.com/google/uuid"
)

type CreateRouteDto struct {
	Name             string    `json:"name" validate:"required"`
	OperationID      uuid.UUID `json:"operation_id" validate:"required"`
	OriginFacilityID uuid.UUID `json:"origin_facility_id" validate:"required"`
	VehicleID        uuid.UUID `json:"vehicle_id" validate:"required"`
}
