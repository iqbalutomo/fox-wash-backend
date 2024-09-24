package dto

type NewOrderRequest struct {
	Address               string                          `json:"address" validate:"required"`
	WashPackageItems      []NewOrderWashPackageItems      `json:"wash_package_items" validate:"required"`
	DetailingPackageItems []NewOrderDetailingPackageItems `json:"detailing_package_items" validate:"required"`
}

type NewOrderWashPackageItems struct {
	WashPackageID uint32 `json:"wash_package_id" validate:"required"`
	Qty           uint32 `json:"qty" validate:"qty"`
}

type NewOrderDetailingPackageItems struct {
	DetailingPackageID uint32 `json:"detailing_package_id" validate:"required"`
	Qty                uint32 `json:"qty" validate:"qty"`
}
