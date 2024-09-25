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

type UpdateWashPackageData struct {
	Name     string  `json:"name" validate:"required"`
	Category uint32  `json:"category" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UpdateWashPackageResponse struct {
	ID        uint32  `json:"id"`
	Name      string  `json:"name"`
	Category  uint32  `json:"category"`
	Price     float64 `json:"price"`
	CreatedBy uint32  `json:"created_by"`
}

type NewDetailingPackageData struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type NewDetailingPackageResponse struct {
	ID          uint32  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedBy   uint32  `json:"created_by"`
}

type UpdateDetailingPackageData struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type UpdateDetailingPackageResponse struct {
	ID          uint32  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedBy   uint32  `json:"created_by"`
}
