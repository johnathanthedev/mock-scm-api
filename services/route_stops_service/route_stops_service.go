package route_stops_service

import (
	"scm-api/api/models"
	"scm-api/db"
	route_stop_dtos "scm-api/types/route-stops/dtos"

	"github.com/google/uuid"
)

func CreateRouteStop(routeStopDto route_stop_dtos.CreateRouteStopDto) (*models.RouteStop, error) {
	newRouteStop := &models.RouteStop{
		RouteID:    routeStopDto.RouteID,
		FacilityID: routeStopDto.FacilityID,
		Sequence:   routeStopDto.Sequence,
	}

	if err := db.GetDB().Create(newRouteStop).Error; err != nil {
		return nil, err
	}

	return newRouteStop, nil
}

func GetRouteStopsByRouteID(routeID uuid.UUID) ([]models.RouteStop, error) {
	var stops []models.RouteStop

	if err := db.GetDB().Where("route_id = ?", routeID).Order("sequence asc").Find(&stops).Error; err != nil {
		return nil, err
	}

	return stops, nil
}
