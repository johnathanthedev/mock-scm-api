package operation_requests

import (
	types "scm-api/types/operations"
)

type CreateOperationRequest struct {
	Name   string                `json:"name" validate:"required,min=3,max=50"`
	Status types.OperationStatus `json:"status" validate:"required,min=3,max=20,oneof=Active Inactive"`
}
