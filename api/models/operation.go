package models

import (
	"time"

	"github.com/google/uuid"
)

type Operation struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name     string    `gorm:"not null"`
	Status   string    `gorm:"not null"`
	Vehicles []Vehicle
	Users    []User `gorm:"many2many:operation_users;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
