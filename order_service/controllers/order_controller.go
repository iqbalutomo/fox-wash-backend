package controllers

import (
	"context"
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
	repo               repository.Order
	userService        userpb.UserClient
	washstationService washstationpb.WashStationClient
	paymentService     services.PaymentService
}

func NewOrderController(repo repository.Order, uc userpb.UserClient, ws washstationpb.WashStationClient, ps services.PaymentService) *OrderController {
	return &OrderController{
		repo:               repo,
		userService:        uc,
		washstationService: ws,
		paymentService:     ps,
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

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	ctxWithAuth := grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	orderItems, err := o.CalculateOrder(ctxWithAuth, washPackageItems)
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

	paymentData, err := o.paymentService.CreateInvoice(newObjectID, orderItems.TotalPrice)
	if err != nil {
		return nil, err
	}

	orderData := models.Order{
		ID: newObjectID,
		OrderDetail: models.OrderDetail{
			WashPackage:      orderItems.WashPackageItems,
			DetailingPackage: nil,
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
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := o.repo.CreateOrder(ctx, &orderData); err != nil {
		return nil, err
	}

	response := helpers.ConvertOrderToResponsePb(orderData)

	return response, nil
}

func (o *OrderController) CalculateOrder(ctx context.Context, items []*orderpb.WashPackageItem) (dto.OrderCalculateResponse, error) {
	fee, err := strconv.ParseFloat(os.Getenv("APP_FEE"), 32)
	if err != nil {
		return dto.OrderCalculateResponse{}, err
	}

	var (
		washPackageIDs       []uint32
		washPackageIDWithQty = map[uint32]uint32{}
		totalPrice           float32
		appFee               float32 = float32(fee)
	)

	for _, item := range items {
		washPackageIDWithQty[item.Id] = item.Qty
		washPackageIDs = append(washPackageIDs, item.Id)
	}

	g, _ := errgroup.WithContext(ctx)
	washPackageDatasChan := make(chan []*models.WashPackage, 1)

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

	if err := g.Wait(); err != nil {
		return dto.OrderCalculateResponse{}, err
	}

	close(washPackageDatasChan)

	var washPackageDatas []models.WashPackage
	for _, wp := range <-washPackageDatasChan {
		qty := washPackageIDWithQty[uint32(wp.ID)]
		totalPrice += wp.Price * float32(qty)

		wpData := models.WashPackage{
			ID:       wp.ID,
			Name:     wp.Name,
			Category: wp.Category,
			Price:    wp.Price,
			Qty:      wp.Qty,
		}

		washPackageDatas = append(washPackageDatas, wpData)
	}

	totalPrice += appFee

	response := dto.OrderCalculateResponse{
		WashPackageItems: washPackageDatas,
		AppFee:           appFee,
		TotalPrice:       totalPrice,
	}

	return response, nil
}
