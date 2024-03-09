package requests

import (
	"github.com/google/uuid"
)

type JoinOperationRequest struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
	UserID      uuid.UUID `json:"user_id" validate:"required"`
}
