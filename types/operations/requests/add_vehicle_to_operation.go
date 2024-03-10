package operation_requests

import "github.com/google/uuid"

type AddVehicleToOperationRequest struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
	VehicleID   uuid.UUID `json:"vehicle_id" validate:"required"`
}
