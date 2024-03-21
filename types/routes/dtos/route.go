package route_dtos

import (
	"time"

	route_stop_dtos "scm-api/types/route-stops/dtos"

	"github.com/google/uuid"
)

type RouteDto struct {
	ID               uuid.UUID                      `json:"id"`
	Name             string                         `json:"name"`
	OperationID      uuid.UUID                      `json:"operation_id"`
	OriginFacilityID uuid.UUID                      `json:"origin_facility_id"`
	VehicleID        uuid.UUID                      `json:"vehicle_id"`
	RouteStops       []route_stop_dtos.RouteStopDto `json:"route_stops"`
	CreatedAt        time.Time                      `json:"created_at"`
	UpdatedAt        time.Time                      `json:"updated_at"`
}
