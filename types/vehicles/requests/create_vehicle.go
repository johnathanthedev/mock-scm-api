package requests

import (
	"encoding/json"
	geo_models "scm-api/api/models/geo"
	vehicle_types "scm-api/types/vehicles"

	"github.com/google/uuid"
)

type CreateVehicleRequest struct {
	VehicleType    vehicle_types.VehicleType   `json:"vehicle_type" validate:"required,oneof=Aircraft"`
	Name           string                      `json:"name" validate:"required,min=3,max=255"`
	CarryVolume    float32                     `json:"carry_volume" validate:"required"`
	MaxWeight      int                         `json:"max_weight" validate:"required"`
	LastLocation   geo_models.GeoPoint         `json:"last_location"`
	Status         vehicle_types.VehicleStatus `json:"status" validate:"required,min=3,max=20,oneof=Active Inactive"`
	PreferredSpeed *float32                    `json:"preferred_speed"`
	CrewCapacity   int                         `json:"crew_capacity" validate:"required"`
	Attributes     *json.RawMessage            `json:"attributes"`
	Make           *string                     `json:"make"`
	Model          *string                     `json:"model"`
	OperationID    uuid.UUID                   `json:"operation_id" validate:"required"`
	DriverID       *uuid.UUID                  `json:"driver_id"`
}
