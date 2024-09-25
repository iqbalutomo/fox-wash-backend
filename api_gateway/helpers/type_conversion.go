package helpers

import (
	"api_gateway/dto"
	"api_gateway/pb/orderpb"
)

func AssertToPbWashPackageItems(orders dto.NewOrderRequest) []*orderpb.WashPackageItem {
	var pbWashPackageItems []*orderpb.WashPackageItem
	for _, wp := range orders.WashPackageItems {
		pbWashPackageItem := &orderpb.WashPackageItem{
			Id:  wp.WashPackageID,
			Qty: wp.Qty,
		}

		pbWashPackageItems = append(pbWashPackageItems, pbWashPackageItem)
	}

	return pbWashPackageItems
}

func AssertToPbDetailingPackageItems(orders dto.NewOrderRequest) []*orderpb.DetailingPackageItem {
	var pbDetailingPackageItems []*orderpb.DetailingPackageItem
	for _, wp := range orders.DetailingPackageItems {
		pbDetailingPackageItem := &orderpb.DetailingPackageItem{
			Id:  wp.DetailingPackageID,
			Qty: wp.Qty,
		}

		pbDetailingPackageItems = append(pbDetailingPackageItems, pbDetailingPackageItem)
	}

	return pbDetailingPackageItems
}
