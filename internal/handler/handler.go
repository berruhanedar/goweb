package handler

import (
	"net/http"

	"github.com/berruhanedar/goweb/internal/dto"
	"github.com/berruhanedar/goweb/internal/errors"
	"github.com/berruhanedar/goweb/internal/mapper"
	"github.com/berruhanedar/goweb/internal/service"
	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	products, err := service.GetAllProducts()
	if err != nil {
		return errors.Respond(c, err, "Failed to retrieve products")
	}
	return c.JSON(http.StatusOK, mapper.ToResponseProducts(products))
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := service.GetProductByID(id)
	if err != nil {
		return errors.Respond(c, err, "Failed to fetch product")
	}
	return c.JSON(http.StatusOK, mapper.ToResponseProduct(*product))
}

func CreateProduct(c echo.Context) error {
	var req dto.CreateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	product := mapper.ToProduct(req)
	newProduct, err := service.CreateProduct(product)
	if err != nil {
		return errors.Respond(c, err, "Failed to create product")
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Product created successfully",
		"product": mapper.ToResponseProduct(*newProduct),
	})
}

func UpdateProduct(c echo.Context) error {
	var req dto.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updateData := mapper.ToUpdateMapFromUpdateRequest(req)
	if len(updateData) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Nothing to update"})
	}

	id := c.Param("id")
	if err := service.UpdateProduct(id, updateData); err != nil {
		return errors.Respond(c, err, "Failed to update product")
	}

	updatedProduct, err := service.GetProductByID(id)
	if err != nil {
		return errors.Respond(c, err, "Product updated but failed to fetch")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Product updated successfully",
		"product": mapper.ToResponseProduct(*updatedProduct),
	})
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	if err := service.DeleteProduct(id); err != nil {
		return errors.Respond(c, err, "Failed to delete product")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Product deleted successfully",
	})
}

func PatchProduct(c echo.Context) error {
	var req dto.PatchProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid body"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	updateData := mapper.ToPatchMap(req)
	if len(updateData) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Nothing to patch"})
	}

	id := c.Param("id")
	if err := service.PatchProduct(id, updateData); err != nil {
		return errors.Respond(c, err, "Failed to patch product")
	}

	patchedProduct, err := service.GetProductByID(id)
	if err != nil {
		return errors.Respond(c, err, "Product patched but failed to fetch")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Product patched successfully",
		"product": mapper.ToResponseProduct(*patchedProduct),
	})
}
