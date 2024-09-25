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
			DetailingPackages: nil,
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
