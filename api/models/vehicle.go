package models

import (
	"encoding/json"
	"time"

	geo_models "scm-api/api/models/geo"
	vehicle_types "scm-api/types/vehicles"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID             uuid.UUID                   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name           string                      `gorm:"not null"`
	Make           *string                     `gorm:"null"`
	Model          *string                     `gorm:"null"`
	Status         vehicle_types.VehicleStatus `gorm:"not null"`
	CrewCapacity   int                         `gorm:"not null"`
	Attributes     *json.RawMessage            `gorm:"type:jsonb;null"`
	PreferredSpeed *float32                    `gorm:"null"`
	VehicleType    vehicle_types.VehicleType   `gorm:"not null"`
	CarryVolume    float32                     `gorm:"not null"`
	MaxWeight      int                         `gorm:"not null"`
	LastLocation   geo_models.GeoPoint         `gorm:"type:geography(Point,4326);null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OperationID    uuid.UUID  `gorm:"not null"`
	DriverID       *uuid.UUID `gorm:"null"`

	Operation Operation     `gorm:"foreignKey:OperationID"`
	Driver    OperationUser `gorm:"foreignKey:DriverID"`
}
