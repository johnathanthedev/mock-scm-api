package models

import (
	"time"

	"github.com/google/uuid"
)

type OperationUser struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OperationID uuid.UUID `gorm:"not null"`
	UserID      uuid.UUID `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Operation Operation `gorm:"foreignKey:OperationID"`
	User      User      `gorm:"foreignKey:UserID"`
}
