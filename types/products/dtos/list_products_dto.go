package product_dtos

import (
	"github.com/google/uuid"
)

type ListProductsDto struct {
	OperationID uuid.UUID `json:"operation_id" validate:"required"`
}
