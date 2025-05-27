package dto

type PatchProductRequest struct {
	Name  *string  `json:"name" validate:"omitempty,min=3"`
	Price *float64 `json:"price" validate:"omitempty,gt=0"`
}
