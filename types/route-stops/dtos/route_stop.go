package route_stop_dtos

import (
	"time"

	geo_models "scm-api/api/models/geo"

	"github.com/google/uuid"
)

type RouteStopDto struct {
	ID                  uuid.UUID           `json:"id"`
	RouteID             uuid.UUID           `json:"route_id"`
	FacilityCoordinates geo_models.GeoPoint `json:"facility_coordinates"`
	Sequence            int                 `json:"sequence"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
}
