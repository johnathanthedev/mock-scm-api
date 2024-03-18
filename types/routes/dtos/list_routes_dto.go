package route_dtos

import (
	"github.com/google/uuid"
)

type ListRoutesDto struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
}
