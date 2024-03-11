package controllers

import (
	"log"
	"net/http"
	"scm-api/services/vehicles_service"
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

	roomID := req.RoomID                                               // This should be dynamic based on your application logic
	locationData := map[string]string{"location": "new location data"} // Example data structure

	trackingService := vehicles_service.NewTrackingService(vc.Broker)
	trackingService.BroadcastLocationUpdate(roomID, locationData)

	log.Println("Location update broadcasted")

	return c.JSON(http.StatusOK, map[string]string{"message": "Location updated"})
}
