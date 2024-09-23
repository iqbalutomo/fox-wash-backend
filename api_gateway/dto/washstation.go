package dto

type NewWashPackageData struct {
	Name     string  `json:"name" validate:"required"`
	Category uint32  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type NewWashPackageResponse struct {
	ID        uint32  `json:"id"`
	Name      string  `json:"name"`
	Category  uint32  `json:"category"`
	Price     float64 `json:"price"`
	CreatedBy uint32  `json:"created_by"`
}
