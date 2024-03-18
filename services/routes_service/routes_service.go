package routes_service

import (
	"errors"
	"scm-api/api/models"
	"scm-api/db"
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

func GetAllRoutesByOperationID(operationID uuid.UUID) ([]models.Route, error) {
	var routes []models.Route

	if err := db.GetDB().Where("operation_id = ?", operationID).Find(&routes).Error; err != nil {
		return nil, err
	}

	return routes, nil
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
