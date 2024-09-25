package dto

import "order_service/models"

type OrderCalculateResponse struct {
	WashPackageItems []models.WashPackage `bson:"wash_package_items" json:"wash_package_items"`
	AppFee           float32              `bson:"app_fee" json:"app_fee"`
	TotalPrice       float32              `bson:"total_price" json:"total_price"`
}
