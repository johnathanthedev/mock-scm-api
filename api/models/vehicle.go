package models

import (
	"time"

	models "scm-api/api/models/geo"

	"github.com/google/uuid"
)

type VehicleType string
type OperationStatus string

const (
	Aircraft VehicleType = "Aircraft"
)

const (
	Active   OperationStatus = "Active"
	Inactive OperationStatus = "Inactive"
)

type Vehicle struct {
	ID                uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OperationID       uuid.UUID       `gorm:"null"`
	VehicleType       VehicleType     `gorm:"not null"`
	Name              string          `gorm:"not null"`
	CarryVolume       float32         `gorm:"not null"`
	MaxWeight         int             `gorm:"not null"`
	LastLocation      models.GeoPoint `gorm:"type:geography(Point,4326)"`
	DepartureLocation models.GeoPoint `gorm:"type:geography(Point,4326)"`
	ArrivalLocation   models.GeoPoint `gorm:"type:geography(Point,4326)"`
	Status            OperationStatus `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
