package controllers

import (
	"log"
	"net/http"
	"scm-api/services/vehicles_service"
	vehicle_dtos "scm-api/types/vehicles/dtos"
	vehicle_requests "scm-api/types/vehicles/requests"

	ws "scm-api/ws"

	"github.com/labstack/echo/v4"
)

type VehicleController struct {
	Broker *ws.Broker
}

func NewVehicleTrackingController(broker *ws.Broker) *VehicleController {
	return &VehicleController{Broker: broker}
}

func (vc *VehicleController) UpdateVehicleLocation(c echo.Context) error {
	req := c.Get("validatedRequest").(*vehicle_requests.UpdateVehicleLocationRequest)

	roomID := req.RoomID

	locationData := vehicle_dtos.LocationDataDto{
		Location: req.Location,
	}

	trackingService := vehicles_service.NewTrackingService(vc.Broker)
	trackingService.BroadcastLocationUpdate(roomID, locationData)

	log.Println("Location update broadcasted")

	return c.JSON(http.StatusOK, map[string]string{"message": "Location updated"})
}
