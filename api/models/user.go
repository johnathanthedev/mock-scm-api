package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username   string      `gorm:"unique;not null"`
	Operations []Operation `gorm:"many2many:operation_users;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
