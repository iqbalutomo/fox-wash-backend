package controllers

import (
	"context"
	"math"
	"order_service/dto"
	"order_service/helpers"
	"order_service/models"
	"order_service/pb/orderpb"
	"order_service/pb/userpb"
	"order_service/pb/washstationpb"
	"order_service/repository"
	"order_service/services"
	"order_service/utils"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
	grpcMetadata "google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	orderpb.UnimplementedOrderServiceServer
	repo                 repository.Order
	userService          userpb.UserClient
	washstationService   washstationpb.WashStationClient
	paymentService       services.PaymentService
	messageBrokerService services.MessageBroker
}

func NewOrderController(repo repository.Order, uc userpb.UserClient, ws washstationpb.WashStationClient, ps services.PaymentService, mb services.MessageBroker) *OrderController {
	return &OrderController{
		repo:                 repo,
		userService:          uc,
		washstationService:   ws,
		paymentService:       ps,
		messageBrokerService: mb,
	}
}

func (o *OrderController) CreateOrder(ctx context.Context, data *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	token, err := helpers.SignJWTForGRPC()
	if err != nil {
		return nil, err
	}

	newObjectID := primitive.NewObjectID()

	userData := models.User{
		ID:    int(data.UserId),
		Name:  data.Name,
		Email: data.Email,
	}

	var washPackageItems []*orderpb.WashPackageItem
	for _, wash := range data.WashPackageItems {
		washTmp := &orderpb.WashPackageItem{
			Id:  wash.Id,
			Qty: wash.Qty,
		}

		washPackageItems = append(washPackageItems, washTmp)
	}

	var detailingPackageItems []*orderpb.DetailingPackageItem
	for _, detailing := range data.DetailingPackageItems {
		detailingTmp := &orderpb.DetailingPackageItem{
			Id:  detailing.Id,
			Qty: detailing.Qty,
		}

		detailingPackageItems = append(detailingPackageItems, detailingTmp)
	}

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	orderItems, err := o.CalculateOrder(ctxWithAuth, washPackageItems, detailingPackageItems)
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	ctxWithAuth = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	availableWasher, err := o.userService.GetAvailableWasher(ctxWithAuth, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	washerData := models.Washer{
		ID:     int(availableWasher.Id),
		Name:   availableWasher.Name,
		Status: availableWasher.Status,
	}

	paymentData, err := o.paymentService.CreateInvoice(newObjectID, orderItems.TotalPrice, userData.Email)
	if err != nil {
		return nil, err
	}

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}
	localTime := time.Now().In(location)

	orderData := models.Order{
		ID: newObjectID,
		OrderDetail: models.OrderDetail{
			WashPackage:      orderItems.WashPackageItems,
			DetailingPackage: orderItems.DetailingPackageItems,
			AppFee:           orderItems.AppFee,
			TotalPrice:       orderItems.TotalPrice,
		},
		User:   userData,
		Washer: washerData,
		Address: models.Address{
			Latitude:  data.Address.Latitude,
			Longitude: data.Address.Longitude,
		},
		Payment:   paymentData,
		Status:    utils.OrderStatusPendingPayment,
		CreatedAt: primitive.NewDateTimeFromTime(localTime),
	}

	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := o.repo.CreateOrder(ctx, &orderData); err != nil {
		return nil, err
	}

	var newMbOrderDetail dto.MbOrderDetail
	for _, orderItem := range orderData.OrderDetail.WashPackage {
		newWashPackageData := dto.MbWashPackage{
			ID:       orderItem.ID,
			Name:     orderItem.Name,
			Category: orderItem.Category,
			Price:    helpers.FormatRupiah(orderItem.Price),
			Qty:      orderItem.Qty,
			SubTotal: helpers.FormatRupiah(orderItem.SubTotal),
		}

		newMbOrderDetail.WashPackage = append(newMbOrderDetail.WashPackage, newWashPackageData)
	}
	for _, orderItem := range orderData.OrderDetail.DetailingPackage {
		newDetailingPackageData := dto.MbDetailingPackage{
			ID:          orderItem.ID,
			Name:        orderItem.Name,
			Description: orderItem.Description,
			Price:       helpers.FormatRupiah(orderItem.Price),
			Qty:         orderItem.Qty,
			SubTotal:    helpers.FormatRupiah(orderItem.SubTotal),
		}

		newMbOrderDetail.DetailingPackage = append(newMbOrderDetail.DetailingPackage, newDetailingPackageData)
	}
	newMbOrderDetail.AppFee = helpers.FormatRupiah(orderData.OrderDetail.AppFee)
	newMbOrderDetail.TotalPrice = helpers.FormatRupiah(orderData.OrderDetail.TotalPrice)

	mbOrderData := map[string]interface{}{
		"order_detail": newMbOrderDetail,
		"User":         orderData.User,
		"Washer":       orderData.Washer,
		"Address":      orderData.Address,
		"Payment":      orderData.Payment,
		"Status":       orderData.Status,
	}

	if err := o.messageBrokerService.PublishMessageOrder(mbOrderData); err != nil {
		return nil, err
	}

	response := helpers.ConvertOrderToResponsePb(orderData)

	return response, nil
}

func (o *OrderController) CalculateOrder(ctx context.Context, washPackageItems []*orderpb.WashPackageItem, detailingPackageItems []*orderpb.DetailingPackageItem) (dto.OrderCalculateResponse, error) {
	fee, err := strconv.ParseFloat(os.Getenv("APP_FEE"), 32)
	if err != nil {
		return dto.OrderCalculateResponse{}, err
	}

	var (
		washPackageIDs            []uint32
		washPackageIDWithQty      = map[uint32]uint32{}
		detailingPackageIDs       []uint32
		detailingPackageIDWithQty = map[uint32]uint32{}
		totalPrice                float32
		appFee                    float32 = float32(fee)
	)

	for _, item := range washPackageItems {
		washPackageIDWithQty[item.Id] = item.Qty
		washPackageIDs = append(washPackageIDs, item.Id)
	}
	for _, item := range detailingPackageItems {
		detailingPackageIDWithQty[item.Id] = item.Qty
		detailingPackageIDs = append(detailingPackageIDs, item.Id)
	}

	g, _ := errgroup.WithContext(ctx)
	washPackageDatasChan := make(chan []*models.WashPackage, 1)
	detailingPackageDatasChan := make(chan []*models.DetailingPackage, 1)

	g.Go(func() error {
		washPackageDatas, err := o.washstationService.FindMultipleWashPackages(ctx, &washstationpb.WashPackageIDs{Ids: washPackageIDs})
		if err != nil {
			return err
		}

		var convertedWashPackages []*models.WashPackage
		for _, wp := range washPackageDatas.WashPackages {
			convertedWashPackages = append(convertedWashPackages, &models.WashPackage{
				ID:       int(wp.Id),
				Name:     wp.Name,
				Category: int(wp.Category),
				Price:    wp.Price,
				Qty:      uint(washPackageIDWithQty[wp.Id]),
			})
		}

		washPackageDatasChan <- convertedWashPackages
		return nil
	})

	g.Go(func() error {
		detailingPackageDatas, err := o.washstationService.FindMultipleDetailingPackages(ctx, &washstationpb.DetailingPackageIDs{Ids: detailingPackageIDs})
		if err != nil {
			return err
		}

		var convertedDetailingPackages []*models.DetailingPackage
		for _, dp := range detailingPackageDatas.DetailingPackages {
			convertedDetailingPackages = append(convertedDetailingPackages, &models.DetailingPackage{
				ID:          int(dp.Id),
				Name:        dp.Name,
				Description: dp.Description,
				Price:       dp.Price,
				Qty:         uint(detailingPackageIDWithQty[dp.Id]),
			})
		}

		detailingPackageDatasChan <- convertedDetailingPackages
		return nil
	})

	if err := g.Wait(); err != nil {
		return dto.OrderCalculateResponse{}, err
	}

	close(washPackageDatasChan)
	close(detailingPackageDatasChan)

	var washPackageDatas []models.WashPackage
	for _, wp := range <-washPackageDatasChan {
		quantity := washPackageIDWithQty[uint32(wp.ID)]
		subtotal := math.Round(float64(wp.Price) * float64(quantity))
		totalPrice += float32(subtotal)

		wpData := models.WashPackage{
			ID:       wp.ID,
			Name:     wp.Name,
			Category: wp.Category,
			Price:    wp.Price,
			Qty:      uint(quantity),
			SubTotal: float32(subtotal),
		}

		washPackageDatas = append(washPackageDatas, wpData)
	}

	var detailingPackageDatas []models.DetailingPackage
	for _, dp := range <-detailingPackageDatasChan {
		quantity := detailingPackageIDWithQty[uint32(dp.ID)]
		subtotal := math.Round(float64(dp.Price) * float64(quantity))
		totalPrice += float32(subtotal)

		dpData := models.DetailingPackage{
			ID:          dp.ID,
			Name:        dp.Name,
			Description: dp.Description,
			Price:       dp.Price,
			Qty:         uint(quantity),
			SubTotal:    float32(subtotal),
		}

		detailingPackageDatas = append(detailingPackageDatas, dpData)
	}

	totalPrice += appFee

	response := dto.OrderCalculateResponse{
		WashPackageItems:      washPackageDatas,
		DetailingPackageItems: detailingPackageDatas,
		AppFee:                appFee,
		TotalPrice:            totalPrice,
	}

	return response, nil
}

func (o *OrderController) UpdateOrderPaymentStatus(ctx context.Context, data *orderpb.UpdatePaymentRequest) (*emptypb.Empty, error) {
	if err := o.repo.UpdateOrderPaymentStatus(ctx, data.InvoiceId, data.Status, data.Method, data.CompletedAt); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (o *OrderController) GetOrderByID(ctx context.Context, data *orderpb.OrderID) (*orderpb.Order, error) {
	orderTmp, err := o.repo.FindByID(ctx, data.Id)
	if err != nil {
		return nil, err
	}

	order := helpers.AssertOrderToPb(orderTmp)

	return order, nil
}

func (o *OrderController) GetWasherAllOrders(ctx context.Context, data *orderpb.WasherID) (*orderpb.Orders, error) {
	ordersTmp, err := o.repo.FindWasherAllOrders(ctx, uint(data.Id))
	if err != nil {
		return nil, err
	}

	orders := helpers.AssertOrdersToPb(ordersTmp)

	return &orderpb.Orders{Orders: orders}, nil
}

func (o *OrderController) GetWasherCurrentOrder(ctx context.Context, data *orderpb.WasherID) (*orderpb.Order, error) {
	orderTmp, err := o.repo.FindWasherCurrentOrder(ctx, uint(data.Id))
	if err != nil {
		return nil, err
	}

	order := helpers.AssertOrderToPb(orderTmp)

	return order, nil
}
