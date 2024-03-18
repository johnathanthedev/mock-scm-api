package products_service

import (
	models "scm-api/api/models"
	"scm-api/db"
	product_dtos "scm-api/types/products/dtos"

	"github.com/google/uuid"
)

func CreateProduct(productDto product_dtos.CreateProductDto) (*models.Product, error) {
	newProduct := &models.Product{
		Name:        productDto.Name,
		Price:       productDto.Price,
		WeightKG:    productDto.WeightKG,
		VolumeM3:    productDto.VolumeM3,
		OperationID: productDto.OperationID,
	}

	if err := db.GetDB().Create(newProduct).Error; err != nil {
		return nil, err
	}

	return newProduct, nil
}

func GetAllProductsByOperationID(operationID uuid.UUID) ([]models.Product, error) {
	var products []models.Product

	if err := db.GetDB().Where("operation_id = ?", operationID).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
