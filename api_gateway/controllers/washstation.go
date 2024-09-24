package controllers

import (
	"api_gateway/dto"
	"api_gateway/helpers"
	washstationpb "api_gateway/pb/washstationpb"
	"api_gateway/utils"
	"fmt"
	"net/http"
	"strconv"

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
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
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

func (w *WashStationController) GetWashPackageByID(c echo.Context) error {
	washPackageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	washPackageData, err := w.client.FindWashPackageByID(ctx, &washstationpb.WashPackageID{Id: uint32(washPackageID)})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: fmt.Sprintf("Get wash package by %d", washPackageID),
		Data:    washPackageData,
	})
}

func (w *WashStationController) UpdateWashPackage(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.AdminRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	var washPackageUpdate dto.UpdateWashPackageData
	washPackageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Bind(&washPackageUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&washPackageUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbUpdateWashPackage := &washstationpb.UpdateWashPackageData{
		Id:       uint32(washPackageID),
		Name:     washPackageUpdate.Name,
		Category: washPackageUpdate.Category,
		Price:    float32(washPackageUpdate.Price),
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	if _, err := w.client.UpdateWashPackage(ctx, pbUpdateWashPackage); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	resp := dto.UpdateWashPackageResponse{
		ID:        pbUpdateWashPackage.Id,
		Name:      pbUpdateWashPackage.Name,
		Category:  pbUpdateWashPackage.Category,
		Price:     float64(pbUpdateWashPackage.Price),
		CreatedBy: uint32(user.ID),
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Wash package has been updated!",
		Data:    resp,
	})
}

func (w *WashStationController) DeleteWashPackage(c echo.Context) error {
	washPackageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.AdminRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	if _, err := w.client.DeleteWashPackage(ctx, &washstationpb.WashPackageID{Id: uint32(washPackageID)}); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Wash package has been deleted!",
		Data:    fmt.Sprintf("Wash package with ID %d", washPackageID),
	})
}
