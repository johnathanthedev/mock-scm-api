package operation_dtos

import (
	operation_types "scm-api/types/operations"

	"github.com/google/uuid"
)

type OperationDto struct {
	ID     uuid.UUID
	Name   string
	Status operation_types.OperationStatus
	Joined bool
}
