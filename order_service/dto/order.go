package dto

import "order_service/models"

type OrderCalculateResponse struct {
	WashPackageItems []models.WashPackage `bson:"wash_package_items" json:"wash_package_items"`
	AppFee           float32              `bson:"app_fee" json:"app_fee"`
	TotalPrice       float32              `bson:"total_price" json:"total_price"`
}

type MbOrderDetail struct {
	WashPackage      []MbWashPackage      `json:"wash_packages"`
	DetailingPackage []MbDetailingPackage `json:"detailing_packages"`
	AppFee           string               `json:"app_fee"`
	TotalPrice       string               `json:"total_price"`
}

type MbWashPackage struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category int    `json:"category"`
	Price    string `json:"price"`
	Qty      uint   `json:"qty"`
	SubTotal string `json:"subtotal"`
}

type MbDetailingPackage struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Qty         uint   `json:"qty"`
	SubTotal    string `json:"subtotal"`
}
