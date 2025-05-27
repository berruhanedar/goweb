package mapper

import (
	"github.com/berruhanedar/goweb/internal/dto"
	"github.com/berruhanedar/goweb/internal/model"
)

func ToProduct(dto dto.CreateProductRequest) model.Product {
	return model.Product{
		Name:  dto.Name,
		Price: dto.Price,
	}
}

func ToUpdateMapFromUpdateRequest(dto dto.UpdateProductRequest) map[string]interface{} {
	return map[string]interface{}{
		"name":  dto.Name,
		"price": dto.Price,
	}
}

func ToPatchMap(dto dto.PatchProductRequest) map[string]interface{} {
	updateData := make(map[string]interface{})
	if dto.Name != nil {
		updateData["name"] = *dto.Name
	}
	if dto.Price != nil {
		updateData["price"] = *dto.Price
	}
	return updateData
}

func ToResponseProduct(model model.Product) dto.ResponseProduct {
	return dto.ResponseProduct{
		ID:    model.ID,
		Name:  model.Name,
		Price: model.Price,
	}
}

func ToResponseProducts(model []model.Product) []dto.ResponseProduct {
	response := make([]dto.ResponseProduct, len(model))
	for i, p := range model {
		response[i] = ToResponseProduct(p)
	}
	return response
}
