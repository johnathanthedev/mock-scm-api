package route_dtos

import (
	"time"

	geo_models "scm-api/api/models/geo"
	route_stop_dtos "scm-api/types/route-stops/dtos"

	"github.com/google/uuid"
)

type RouteDto struct {
	ID                  uuid.UUID                      `json:"id"`
	Name                string                         `json:"name"`
	OperationID         uuid.UUID                      `json:"operation_id"`
	FacilityCoordinates geo_models.GeoPoint            `json:"facility_coordinates"`
	VehicleID           uuid.UUID                      `json:"vehicle_id"`
	RouteStops          []route_stop_dtos.RouteStopDto `json:"route_stops"`
	CreatedAt           time.Time                      `json:"created_at"`
	UpdatedAt           time.Time                      `json:"updated_at"`
}
