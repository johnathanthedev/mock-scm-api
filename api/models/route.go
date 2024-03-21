package models

import (
	"time"

	"github.com/google/uuid"
)

type Route struct {
	ID               uuid.UUID   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name             string      `gorm:"type:varchar(255);not null"`
	OperationID      uuid.UUID   `gorm:"not null"`
	OriginFacilityID uuid.UUID   `gorm:"not null"`
	VehicleID        uuid.UUID   `gorm:"not null"`
	Operation        Operation   `gorm:"foreignKey:OperationID"`
	OriginFacility   Facility    `gorm:"foreignKey:OriginFacilityID"`
	Vehicle          Vehicle     `gorm:"foreignKey:VehicleID"`
	RouteStops       []RouteStop `gorm:"foreignKey:RouteID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
