package requests

import (
	"github.com/google/uuid"
)

type LeaveOperationRequest struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
}
