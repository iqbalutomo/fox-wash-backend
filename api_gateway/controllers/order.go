package controllers

import (
	"api_gateway/dto"
	"api_gateway/helpers"
	"api_gateway/pb/orderpb"
	"api_gateway/pb/userpb"
	"api_gateway/services"
	"api_gateway/utils"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	client     orderpb.OrderServiceClient
	userClient userpb.UserClient
	maps       services.Maps
}

func NewOrderController(client orderpb.OrderServiceClient, userClient userpb.UserClient, maps services.Maps) *OrderController {
	return &OrderController{client, userClient, maps}
}

func (o *OrderController) CreateOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.UserRole {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails("Access permission"))
	}

	var orderRequest dto.NewOrderRequest
	if err := c.Bind(&orderRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&orderRequest); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	coordinate, err := o.maps.GetCoordinate(orderRequest.Address)
	if err != nil {
		return err
	}

	pbWashPackageItems := helpers.AssertToPbWashPackageItems(orderRequest)
	pbDetailingPackageItems := helpers.AssertToPbDetailingPackageItems(orderRequest)

	pbOrderRequest := &orderpb.CreateOrderRequest{
		UserId: uint32(user.ID),
		Name:   user.Name,
		Email:  user.Email,
		Address: &orderpb.Address{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		},
		WashPackageItems:      pbWashPackageItems,
		DetailingPackageItems: pbDetailingPackageItems,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	res, err := o.client.CreateOrder(ctx, pbOrderRequest)
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Order has been created",
		Data:    res,
	})
}

func (o *OrderController) UpdatePaymentStatus(c echo.Context) error {
	webhookToken := c.Request().Header.Get("x-callback-token")
	if webhookToken != os.Getenv("XENDIT_WEBHOOK_TOKEN") {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid webhook token"))
	}

	var paymentData dto.InvoiceXenditWebhook
	if err := c.Bind(&paymentData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbPaymentData := &orderpb.UpdatePaymentRequest{
		InvoiceId:   paymentData.InvoiceID,
		Status:      paymentData.Status,
		Method:      paymentData.PaymentMethod,
		CompletedAt: paymentData.PaidAt,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	if _, err := o.client.UpdateOrderPaymentStatus(ctx, pbPaymentData); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	ctx, cancel, err = helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	mbPaymentData := &userpb.PaymentSuccessData{
		InvoiceId:   pbPaymentData.InvoiceId,
		Status:      pbPaymentData.Status,
		Method:      pbPaymentData.Method,
		CompletedAt: pbPaymentData.CompletedAt,
		PayerEmail:  paymentData.PayerEmail,
	}

	_, err = o.userClient.PostPublishMessagePaymentSuccess(ctx, mbPaymentData)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
