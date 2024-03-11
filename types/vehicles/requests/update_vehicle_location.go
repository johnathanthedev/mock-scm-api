package requests

type UpdateVehicleLocationRequest struct {
	RoomID string `json:"room_id" validate:"required"`
}
