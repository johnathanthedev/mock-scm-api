package models

import (
	"time"

	operation_types "scm-api/types/operations"

	"github.com/google/uuid"
)

type Operation struct {
	ID     uuid.UUID                       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name   string                          `gorm:"unique,not null"`
	Status operation_types.OperationStatus `gorm:"not null"`
	Users  []User                          `gorm:"many2many:operation_users;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
