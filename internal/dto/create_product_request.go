package dto

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required,min=3"`
	Price float64 `json:"price" validate:"required,gt=0"`
}
