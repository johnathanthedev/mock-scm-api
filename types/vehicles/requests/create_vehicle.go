package requests

import (
	"encoding/json"
	models "scm-api/api/models/geo"
	vehicle_types "scm-api/types/vehicles"

	"github.com/google/uuid"
)

type CreateVehicleRequest struct {
	OperationID       *uuid.UUID                  `json:"operation_id"`
	VehicleType       vehicle_types.VehicleType   `json:"vehicle_type" validate:"required, oneof=Aircraft"`
	Name              string                      `json:"name" validate:"required,min=3,max=50"`
	CarryVolume       float32                     `json:"carry_volume" validate:"required"`
	MaxWeight         int                         `json:"max_weight" validate:"required"`
	LastLocation      models.GeoPoint             `json:"last_location"`
	DepartureLocation models.GeoPoint             `json:"departure_location"`
	ArrivalLocation   models.GeoPoint             `json:"arrival_location"`
	Status            vehicle_types.VehicleStatus `json:"status" validate:"required, min=3, max=20, oneof=Active Inactive"`
	PreferredSpeed    int                         `json:"preferred_speed" validate:"required"`
	CrewCapacity      int                         `json:"crew_capacity" validate:"required"`
	Attributes        json.RawMessage             `json:"attributes"`
	Make              string                      `json:"make"`
	Model             string                      `json:"model"`
}
