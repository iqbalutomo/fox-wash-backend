package dto

import (
	"api_gateway/models"
	"api_gateway/pb/orderpb"
	"api_gateway/pb/washstationpb"
)

type SwaggerResponseRegister struct {
	Message string      `json:"message" extensions:"x-order=0"`
	Data    models.User `json:"data" extensions:"x-order=1"`
}

type SwaggerNewWashPackageData struct {
	Name     string  `json:"name" validate:"required" extensions:"x-order=0"`
	Category uint32  `json:"category" validate:"required" extensions:"x-order=1"`
	Price    float64 `json:"price" validate:"required" extensions:"x-order=2"`
}

type SwaggerResponseNewWashPackageByAdmin struct {
	Message string                 `json:"message" extensions:"x-order=0"`
	Data    NewWashPackageResponse `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseGetAllWashPackage struct {
	Message string                                    `json:"message" extensions:"x-order=0"`
	Data    *washstationpb.WashPackageCompactRepeated `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseGetWashPackageByID struct {
	Message string                            `json:"message" extensions:"x-order=0"`
	Data    *washstationpb.WashPackageCompact `json:"data" extensions:"x-order=1"`
}

type SwaggerUpdateWashPackageData struct {
	Name     string  `json:"name" validate:"required" extensions:"x-order=0"`
	Category uint32  `json:"category" validate:"required" extensions:"x-order=1"`
	Price    float64 `json:"price" validate:"required" extensions:"x-order=2"`
}

type SwaggerResponseOrder struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Order `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseWasherGetAllOrders struct {
	Message string          `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Orders `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseWasherGetCurrentOrder struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Order `json:"data" extensions:"x-order=1"`
}

type SwaggerResponseUpdateOrderStatus struct {
	Message string         `json:"message" extensions:"x-order=0"`
	Data    *orderpb.Order `json:"data" extensions:"x-order=1"`
}
