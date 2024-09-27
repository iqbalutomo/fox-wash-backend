package dto

type NewWashPackageData struct {
	Name     string  `json:"name" validate:"required" extensions:"x-order=0"`
	Category uint32  `json:"category" validate:"required" extensions:"x-order=1"`
	Price    float64 `json:"price" validate:"required" extensions:"x-order=2"`
}

type NewWashPackageResponse struct {
	ID        uint32  `json:"id" extensions:"x-order=0"`
	Name      string  `json:"name" extensions:"x-order=1"`
	Category  uint32  `json:"category" extensions:"x-order=2"`
	Price     float64 `json:"price" extensions:"x-order=3"`
	CreatedBy uint32  `json:"created_by" extensions:"x-order=4"`
}

type UpdateWashPackageData struct {
	Name     string  `json:"name" validate:"required" extensions:"x-order=0"`
	Category uint32  `json:"category" validate:"required" extensions:"x-order=1"`
	Price    float64 `json:"price" validate:"required" extensions:"x-order=2"`
}

type UpdateWashPackageResponse struct {
	ID        uint32  `json:"id" extensions:"x-order=0"`
	Name      string  `json:"name" extensions:"x-order=1"`
	Category  uint32  `json:"category" extensions:"x-order=2"`
	Price     float64 `json:"price" extensions:"x-order=3"`
	CreatedBy uint32  `json:"created_by" extensions:"x-order=4"`
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
