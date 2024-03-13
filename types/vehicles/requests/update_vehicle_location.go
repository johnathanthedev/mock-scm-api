package requests

import models "scm-api/api/models/geo"

type UpdateVehicleLocationRequest struct {
	RoomID   string          `json:"room_id" validate:"required"`
	Location models.GeoPoint `json:"location" validate:"required"`
}
