package controllers

import (
	"api_gateway/dto"
	"api_gateway/helpers"
	washstationpb "api_gateway/pb/washstationpb"
	"api_gateway/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WashStationController struct {
	client washstationpb.WashStationClient
}

func NewWashStationController(client washstationpb.WashStationClient) *WashStationController {
	return &WashStationController{client}
}

func (w *WashStationController) CreateWashPackage(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.AdminRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	var washPackageData dto.NewWashPackageData
	if err := c.Bind(&washPackageData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&washPackageData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	newWashPackageData := &washstationpb.NewWashPackageData{
		Name:      washPackageData.Name,
		Category:  washPackageData.Category,
		Price:     float32(washPackageData.Price),
		CreatedBy: uint32(user.ID),
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	data, err := w.client.CreateWashPackage(ctx, newWashPackageData)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	response := dto.NewWashPackageResponse{
		ID:        data.Id,
		Name:      newWashPackageData.Name,
		Category:  newWashPackageData.Category,
		Price:     float64(newWashPackageData.Price),
		CreatedBy: newWashPackageData.CreatedBy,
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Wash package has been created!",
		Data:    response,
	})
}

func (w *WashStationController) GetAllWashPackages(c echo.Context) error {
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	washPackageData, err := w.client.FindAllWashPackages(ctx, &emptypb.Empty{})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all wash packages",
		Data:    washPackageData,
	})
}
