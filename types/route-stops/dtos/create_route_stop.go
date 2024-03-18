package route_stop_dtos

import "github.com/google/uuid"

type CreateRouteStopDto struct {
	RouteID    uuid.UUID `json:"route_id" validate:"required"`
	FacilityID uuid.UUID `json:"facility_id" validate:"required"`
	Sequence   int       `json:"sequence" validate:"required,gte=1"`
}
