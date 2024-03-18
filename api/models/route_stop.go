package models

import (
	"time"

	"github.com/google/uuid"
)

type RouteStop struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RouteID    uuid.UUID `gorm:"not null"`
	FacilityID uuid.UUID `gorm:"not null"`
	Sequence   int       `gorm:"default:1;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Route    Route    `gorm:"foreignKey:RouteID"`
	Facility Facility `gorm:"foreignKey:FacilityID"`
}
