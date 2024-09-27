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

// Admin   godoc
// @Summary      Create wash package for admin
// @Description  Creates new package data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         admin
// @Accept       json
// @Produce      json
// @param        request body dto.SwaggerNewWashPackageData true "Create Wash Package"
// @Success      201  {object}  dto.SwaggerResponseNewWashPackageByAdmin
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washstations/wash-package [post]
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

// Admin      godoc
// @Summary      Get all wash package datas
// @Description  Retrieve all wash package datas from the database.
// @Tags         all user
// @Produce      json
// @Success      200  {object}  dto.SwaggerResponseGetAllWashPackage
// @Failure      400  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washstations/wash-package/all [get]
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

// Admin      godoc
// @Summary      Get wash package by ID
// @Description  Retrieve specific wash package data using the wash package id.
// @Tags         all user
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.SwaggerResponseGetWashPackageByID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washstations/wash-package/{id} [get]
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

// Admin      godoc
// @Summary      Update wash package by admin
// @Description  Updates existing package data specific to the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @param 		request body dto.SwaggerUpdateWashPackageData  true  "Update Wash Package"
// @Success      200  {object}  dto.SwaggerResponseGetWashPackageByID
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washstations/wash-package/{id} [put]
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

// Admin      godoc
// @Summary      Delete wash package by admin
// @Description  Deletes existing package for the current logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param 		 id   path      int  true  "Id"
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  utils.ErrResponse
// @Failure      401  {object}  utils.ErrResponse
// @Failure      404  {object}  utils.ErrResponse
// @Failure      500  {object}  utils.ErrResponse
// @Router       /washstations/wash-package/{id} [delete]
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

func (w *WashStationController) CreateDetailingPackage(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.AdminRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	var detailingPackageData dto.NewDetailingPackageData
	if err := c.Bind(&detailingPackageData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&detailingPackageData); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	newDetailingPackageData := &washstationpb.NewDetailingPackageData{
		Name:        detailingPackageData.Name,
		Description: detailingPackageData.Description,
		Price:       float32(detailingPackageData.Price),
		CreatedBy:   uint32(user.ID),
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	data, err := w.client.CreateDetailingPackage(ctx, newDetailingPackageData)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	response := dto.NewDetailingPackageResponse{
		ID:          data.Id,
		Name:        newDetailingPackageData.Name,
		Description: newDetailingPackageData.Description,
		Price:       float64(newDetailingPackageData.Price),
		CreatedBy:   newDetailingPackageData.CreatedBy,
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Detailing package has been created!",
		Data:    response,
	})
}

func (w *WashStationController) GetAllDetailingPackages(c echo.Context) error {
	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	detailingPackageData, err := w.client.FindAllDetailingPackages(ctx, &emptypb.Empty{})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Get all detailing packages",
		Data:    detailingPackageData,
	})
}

func (w *WashStationController) GetDetailingPackageByID(c echo.Context) error {
	detailingPackageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	detailingPackageData, err := w.client.FindDetailingPackageByID(ctx, &washstationpb.DetailingPackageID{Id: uint32(detailingPackageID)})
	if err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: fmt.Sprintf("Get detailing package by %d", detailingPackageID),
		Data:    detailingPackageData,
	})
}

func (w *WashStationController) UpdateDetailingPackage(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.AdminRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	var detailingPackageUpdate dto.UpdateDetailingPackageData
	detailingPackageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Bind(&detailingPackageUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(&detailingPackageUpdate); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbUpdateDetailingPackage := &washstationpb.UpdateDetailingPackageData{
		Id:          uint32(detailingPackageID),
		Name:        detailingPackageUpdate.Name,
		Description: detailingPackageUpdate.Description,
		Price:       float32(detailingPackageUpdate.Price),
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}
	defer cancel()

	if _, err := w.client.UpdateDetailingPackage(ctx, pbUpdateDetailingPackage); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	resp := dto.UpdateDetailingPackageResponse{
		ID:          pbUpdateDetailingPackage.Id,
		Name:        pbUpdateDetailingPackage.Name,
		Description: pbUpdateDetailingPackage.Description,
		Price:       float64(pbUpdateDetailingPackage.Price),
		CreatedBy:   uint32(user.ID),
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Detailing package has been updated!",
		Data:    resp,
	})
}

func (w *WashStationController) DeleteDetailingPackage(c echo.Context) error {
	detailingPackageID, err := strconv.Atoi(c.Param("id"))
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

	if _, err := w.client.DeleteDetailingPackage(ctx, &washstationpb.DetailingPackageID{Id: uint32(detailingPackageID)}); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Detailing package has been deleted!",
		Data:    fmt.Sprintf("Detailing package with ID %d", detailingPackageID),
	})
}
