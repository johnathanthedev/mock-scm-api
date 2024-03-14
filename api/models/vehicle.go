package models

import (
	"encoding/json"
	"time"

	models "scm-api/api/models/geo"
	vehicle_types "scm-api/types/vehicles"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID             uuid.UUID                   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	VehicleType    vehicle_types.VehicleType   `gorm:"not null"`
	Name           string                      `gorm:"unique;not null"`
	CarryVolume    float32                     `gorm:"not null"`
	MaxWeight      int                         `gorm:"not null"`
	LastLocation   models.GeoPoint             `gorm:"type:geography(Point,4326)"`
	Status         vehicle_types.VehicleStatus `gorm:"not null"`
	PreferredSpeed int                         `gorm:"not null"` // Measured in kilometers (km/h)
	CrewCapacity   int                         `gorm:"not null"`
	Attributes     json.RawMessage             `gorm:"type:jsonb;null"`
	Make           string                      `gorm:"null"`
	Model          string                      `gorm:"null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
