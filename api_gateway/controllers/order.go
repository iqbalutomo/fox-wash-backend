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

// @Summary 	Create new order
// @Description Create a new order for the logged in user. You will need an 'Authorization' cookie attached with this request.
// @Tags 		customer
// @Accept 		json
// @Produce 	json
// @Param 		orderRequest body dto.NewOrderRequest true "Order details"
// @Success 	201 {object} dto.SwaggerResponseOrder
// @Failure 	400 {object} utils.ErrResponse
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/orders [post]
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

// @Summary 	Get all user orders
// @Description Retrieves all orders for the logged-in user. You will need an 'Authorization' cookie attached with this request.
// @Tags 		customer
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.SwaggerResponseUserGetAllOrders
// @Failure 	401 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/orders [get]
func (o *OrderController) GetUserAllOrders(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.UserRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Access permission"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.client.GetUserAllOrders(ctx, &orderpb.WasherID{Id: uint32(user.ID)})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all users orders",
		Data:    orders,
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

// Orders        godoc
// @Summary      Get all washer's orders
// @Description  Retrieve all orders related to the logged in washer. You will need an 'Authorization' cookie attached with this request.
// @Tags         washer
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseWasherGetAllOrders
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washers/orders [get]
func (o *OrderController) GetWasherAllOrders(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.WasherRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Access permission"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	orders, err := o.client.GetWasherAllOrders(ctx, &orderpb.WasherID{Id: uint32(user.ID)})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all washers orders",
		Data:    orders,
	})
}

// @Summary      Get order by ID
// @Description  Retrieve an order by it's ID. You will need an 'Authorization' cookie attached with this request.
// @Tags         washer
// @Accept       json
// @Produce      json
// @Param 		 id   path      string  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseOrder
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washers/orders/{id} [get]
func (o *OrderController) WasherGetOrderByID(c echo.Context) error {
	orderID := c.Param("id")

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.WasherRole {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Access permission"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	order, err := o.client.GetOrderByID(ctx, &orderpb.OrderID{Id: orderID})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get order by ID success",
		Data:    order,
	})
}

// Orders        godoc
// @Summary      Get current washer's order
// @Description  Retrieve ongoing order related to the logged in washer. You will need an 'Authorization' cookie attached with this request.
// @Tags         washer
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseWasherGetCurrentOrder
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washers/orders/ongoing [get]
func (o *OrderController) WasherGetCurrentOrder(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.WasherRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	order, err := o.client.GetWasherCurrentOrder(ctx, &orderpb.WasherID{Id: uint32(user.ID)})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get washers current order success",
		Data:    order,
	})
}

// Orders        godoc
// @Summary      Update washer's order
// @Description  Updates ongoing order status related to the logged in washer. You will need an 'Authorization' cookie attached with this request.
// @Tags         washer
// @Accept       json
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseUpdateOrderStatus
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washers/orders/status/{id} [put]
func (o *OrderController) UpdateWasherOrderStatus(c echo.Context) error {
	orderID := c.Param("id")

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.WasherRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	updateOrderStatusReq := &orderpb.UpdateOrderStatusRequest{
		OrderId:  &orderpb.OrderID{Id: orderID},
		WasherId: &orderpb.WasherID{Id: uint32(user.ID)},
	}

	order, err := o.client.UpdateWasherOrderStatus(ctx, updateOrderStatusReq)
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Order status has been updated",
		Data:    order,
	})
}
