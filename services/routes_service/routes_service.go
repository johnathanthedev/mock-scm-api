package routes_service

import (
	"errors"
	"scm-api/api/models"
	geo_models "scm-api/api/models/geo"
	"scm-api/db"
	route_stop_dtos "scm-api/types/route-stops/dtos"
	route_dtos "scm-api/types/routes/dtos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateRoute(routeDto route_dtos.CreateRouteDto) (*models.Route, error) {
	newRoute := &models.Route{
		Name:             routeDto.Name,
		OperationID:      routeDto.OperationID,
		OriginFacilityID: routeDto.OriginFacilityID,
		VehicleID:        routeDto.VehicleID,
	}

	if err := db.GetDB().Create(newRoute).Error; err != nil {
		return nil, err
	}

	return newRoute, nil
}

func GetAllRoutesByOperationID(operationID uuid.UUID) ([]route_dtos.RouteDto, error) {
	var modelRoutes []models.Route
	var dtoRoutes []route_dtos.RouteDto

	// Preload associated Facilities
	if err := db.GetDB().
		Preload("OriginFacility").
		Preload("RouteStops.Facility").
		Where("operation_id = ?", operationID).
		Find(&modelRoutes).Error; err != nil {
		return nil, err
	}

	for _, modelRoute := range modelRoutes {
		var routeStopDtos []route_stop_dtos.RouteStopDto

		for _, modelRouteStop := range modelRoute.RouteStops {
			routeStopDto := route_stop_dtos.RouteStopDto{
				ID:      modelRouteStop.ID,
				RouteID: modelRouteStop.RouteID,
				FacilityCoordinates: geo_models.GeoPoint{
					Lat: modelRouteStop.Facility.Location.Lat,
					Lng: modelRouteStop.Facility.Location.Lng,
				},
				Sequence:  modelRouteStop.Sequence,
				CreatedAt: modelRouteStop.CreatedAt,
				UpdatedAt: modelRouteStop.UpdatedAt,
			}
			routeStopDtos = append(routeStopDtos, routeStopDto)
		}

		dtoRoute := route_dtos.RouteDto{
			ID:          modelRoute.ID,
			Name:        modelRoute.Name,
			OperationID: modelRoute.OperationID,
			FacilityCoordinates: geo_models.GeoPoint{
				Lat: modelRoute.OriginFacility.Location.Lat,
				Lng: modelRoute.OriginFacility.Location.Lng,
			}, VehicleID: modelRoute.VehicleID,
			RouteStops: routeStopDtos,
			CreatedAt:  modelRoute.CreatedAt,
			UpdatedAt:  modelRoute.UpdatedAt,
		}
		dtoRoutes = append(dtoRoutes, dtoRoute)
	}

	return dtoRoutes, nil
}

func GetRouteByID(routeID uuid.UUID) (*models.Route, error) {
	var route models.Route
	result := db.GetDB().First(&route, "id = ?", routeID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("route not found")
		}
		return nil, result.Error
	}

	return &route, nil
}
