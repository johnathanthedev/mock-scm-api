package operation_dtos

import (
	"github.com/google/uuid"
)

type GetOperationDto struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
}
