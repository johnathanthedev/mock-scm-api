package route_stop_dtos

import "github.com/google/uuid"

type ListRouteStopsDto struct {
	RouteID uuid.UUID `json:"route_id" validate:"required"`
}
