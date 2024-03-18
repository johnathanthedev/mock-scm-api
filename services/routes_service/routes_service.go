package routes_service

import (
	"scm-api/api/models"
	"scm-api/db"
	route_dtos "scm-api/types/routes/dtos"

	"github.com/google/uuid"
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

func GetAllRoutesByOperationID(operationID uuid.UUID) ([]models.Route, error) {
	var routes []models.Route

	if err := db.GetDB().Where("operation_id = ?", operationID).Find(&routes).Error; err != nil {
		return nil, err
	}

	return routes, nil
}
