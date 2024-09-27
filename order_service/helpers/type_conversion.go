package helpers

import (
	"order_service/models"
	"order_service/pb/orderpb"
	"order_service/utils"
)

func ConvertOrderToResponsePb(orderData models.Order) *orderpb.CreateOrderResponse {
	return &orderpb.CreateOrderResponse{
		ObjectId: orderData.ID.Hex(),
		OrderDetail: &orderpb.OrderDetail{
			WashPackages:      convertWashPackages(orderData.OrderDetail.WashPackage),
			DetailingPackages: convertDetailingPackages(orderData.OrderDetail.DetailingPackage),
			AppFee:            orderData.OrderDetail.AppFee,
			TotalPrice:        orderData.OrderDetail.TotalPrice,
		},
		User: &orderpb.User{
			Id:    uint32(orderData.User.ID),
			Name:  orderData.User.Name,
			Email: orderData.User.Email,
		},
		Washer: &orderpb.Washer{
			Id:     uint32(orderData.Washer.ID),
			Name:   orderData.Washer.Name,
			Status: orderData.Washer.Status,
		},
		Address: &orderpb.Address{
			Latitude:  orderData.Address.Latitude,
			Longitude: orderData.Address.Longitude,
		},
		Payment: &orderpb.Payment{
			InvoiceId:  orderData.Payment.InvoiceID,
			InvoiceUrl: orderData.Payment.InvoiceURL,
			Total:      orderData.Payment.Total,
			Method:     orderData.Payment.Method,
			Status:     orderData.Payment.Status,
		},
		Status:    utils.OrderStatusPendingPayment,
		CreatedAt: orderData.CreatedAt.Time().String(),
	}
}

func convertWashPackages(washPackages []models.WashPackage) []*orderpb.WashPackage {
	var result []*orderpb.WashPackage
	for _, wp := range washPackages {
		result = append(result, &orderpb.WashPackage{
			Id:       uint32(wp.ID),
			Name:     wp.Name,
			Category: uint32(wp.Category),
			Price:    wp.Price,
			Qty:      uint32(wp.Qty),
			Subtotal: wp.SubTotal,
		})
	}

	return result
}

func convertDetailingPackages(detailingPackages []models.DetailingPackage) []*orderpb.DetailingPackage {
	var result []*orderpb.DetailingPackage
	for _, dp := range detailingPackages {
		result = append(result, &orderpb.DetailingPackage{
			Id:          uint32(dp.ID),
			Name:        dp.Name,
			Description: dp.Description,
			Price:       dp.Price,
			Qty:         uint32(dp.Qty),
			Subtotal:    dp.SubTotal,
		})
	}

	return result
}

func AssertOrdersToPb(ordersTmp []models.Order) []*orderpb.Order {
	orders := []*orderpb.Order{}
	for _, orderTmp := range ordersTmp {
		order := AssertOrderToPb(orderTmp)
		orders = append(orders, order)
	}

	return orders
}

func AssertOrderToPb(orderTmp models.Order) *orderpb.Order {
	var washPackages []*orderpb.WashPackage
	for _, washPackageTmp := range orderTmp.OrderDetail.WashPackage {
		washPackage := &orderpb.WashPackage{
			Id:       uint32(washPackageTmp.ID),
			Name:     washPackageTmp.Name,
			Category: uint32(washPackageTmp.Category),
			Price:    washPackageTmp.Price,
			Qty:      uint32(washPackageTmp.Qty),
			Subtotal: washPackageTmp.SubTotal,
		}

		washPackages = append(washPackages, washPackage)
	}

	var detailingPackages []*orderpb.DetailingPackage
	for _, detailingPackageTmp := range orderTmp.OrderDetail.DetailingPackage {
		detailingPackage := &orderpb.DetailingPackage{
			Id:          uint32(detailingPackageTmp.ID),
			Name:        detailingPackageTmp.Name,
			Description: detailingPackageTmp.Description,
			Price:       detailingPackageTmp.Price,
			Qty:         uint32(detailingPackageTmp.Qty),
			Subtotal:    detailingPackageTmp.SubTotal,
		}

		detailingPackages = append(detailingPackages, detailingPackage)
	}

	orderDetailData := &orderpb.OrderDetail{
		WashPackages:      washPackages,
		DetailingPackages: detailingPackages,
		AppFee:            orderTmp.OrderDetail.AppFee,
		TotalPrice:        orderTmp.OrderDetail.TotalPrice,
	}

	userData := &orderpb.User{
		Id:    uint32(orderTmp.User.ID),
		Name:  orderTmp.User.Name,
		Email: orderTmp.User.Email,
	}

	washerData := &orderpb.Washer{
		Id:     uint32(orderTmp.Washer.ID),
		Name:   orderTmp.Washer.Name,
		Status: orderTmp.Washer.Status,
	}

	paymentData := &orderpb.Payment{
		InvoiceId:  orderTmp.Payment.InvoiceID,
		InvoiceUrl: orderTmp.Payment.InvoiceURL,
		Total:      orderTmp.Payment.Total,
		Method:     orderTmp.Payment.Method,
		Status:     orderTmp.Payment.Status,
	}

	order := &orderpb.Order{
		ObjectId:    orderTmp.ID.Hex(),
		OrderDetail: orderDetailData,
		User:        userData,
		Washer:      washerData,
		Address: &orderpb.Address{
			Latitude:  orderTmp.Address.Latitude,
			Longitude: orderTmp.Address.Longitude,
		},
		Payment:   paymentData,
		Status:    orderTmp.Status,
		CreatedAt: orderTmp.CreatedAt.Time().String(),
	}

	return order
}
