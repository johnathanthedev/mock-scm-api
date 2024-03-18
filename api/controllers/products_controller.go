package controllers

import (
	"net/http"
	"scm-api/services/operations_service"
	"scm-api/services/products_service"
	product_dtos "scm-api/types/products/dtos"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	req := c.Get("validatedRequest").(*product_dtos.CreateProductDto)

	_, err := operations_service.GetOperationByID(req.OperationID)
	if err != nil {
		if err.Error() == "operation not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
		}
	}

	newProduct, err := products_service.CreateProduct(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Product"})
	}

	return c.JSON(http.StatusCreated, newProduct)
}

func ListProducts(c echo.Context) error {
	req := c.Get("validatedRequest").(*product_dtos.ListProductsDto)

	_, err := operations_service.GetOperationByID(req.OperationID)
	if err != nil {
		if err.Error() == "operation not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
		}
	}

	products, err := products_service.GetAllProductsByOperationID(req.OperationID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve products"})
	}

	return c.JSON(http.StatusCreated, products)

}
