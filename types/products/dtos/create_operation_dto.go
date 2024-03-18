package product_dtos

import (
	"github.com/google/uuid"
)

type CreateProductDto struct {
	Name        string    `json:"name" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	WeightKG    int       `json:"weight_kg" validate:"required"`
	VolumeM3    float64   `json:"volume_m3" validate:"required"`
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
}
