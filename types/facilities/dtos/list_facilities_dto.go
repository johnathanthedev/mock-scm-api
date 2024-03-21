package facility_dtos

import (
	"github.com/google/uuid"
)

type ListFacilitiesDto struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
}
