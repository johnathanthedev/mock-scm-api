package models

import (
	"time"

	geo_models "scm-api/api/models/geo"

	"github.com/google/uuid"
)

type Facility struct {
	ID                 uuid.UUID           `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name               string              `gorm:"type:varchar(255);not null"`
	Type               string              `gorm:"type:varchar(255);not null"`
	MaxStorageCapacity int                 `gorm:"not null"`
	DailyOperatingCost int                 `gorm:"not null"`
	DailyRentCost      int                 `gorm:"not null"`
	DailyCarbonOutput  int                 `gorm:"not null"`
	Location           geo_models.GeoPoint `gorm:"type:geography(Point,4326);not null"`
	OperationID        uuid.UUID           `gorm:"type:uuid;not null"`

	Operation Operation `gorm:"foreignKey:OperationID;references:ID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
