package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `gorm:"size:255;not null"`
	Price       float64   `gorm:"type:numeric(10,2);not null"`
	WeightKG    int       `gorm:"not null"`
	VolumeM3    float64   `gorm:"type:numeric(10,3);not null"`
	OperationID uuid.UUID `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	Operation Operation `gorm:"foreignKey:OperationID"`
}
