package controllers_test

import (
	"context"
	"testing"
	"time"
	"wash_station_service/controllers"
	"wash_station_service/models"
	"wash_station_service/pb"
	"wash_station_service/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	mockRepository        = repository.NewMockWashStationRepository()
	washStationController = controllers.NewWashStationController(&mockRepository)
)

var (
	mockWashPackage = &models.Wash{
		ID:        1,
		Name:      "Wash Package 1",
		Category:  200,
		Price:     100.00,
		CreatedBy: 1,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	mockDetailingPackage = &models.Detailing{
		ID:          1,
		Name:        "Detailing Package 1",
		Description: "Detailing Package 1 Description",
		Price:       200.00,
		CreatedBy:   1,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestFindAllWashPackages(t *testing.T) {
	mockWashPackage := models.Wash{
		Name:     "Wash Package 1",
		Category: 200,
		Price:    100.00,
	}

	mockRepository.Mock.On("FindAllWashPackages").Return([]models.Wash{mockWashPackage}, nil)

	_, err := washStationController.FindAllWashPackages(context.Background(), &emptypb.Empty{})

	assert.Nil(t, err)
	assert.NotEmpty(t, mockWashPackage)
}

func TestCreateWashPackage(t *testing.T) {
	var (
		adminId       uint32 = 1
		washPackageId uint32 = 1
	)

	pbRequest := &pb.NewWashPackageData{
		Name:      "Wash Package 1",
		Category:  200,
		Price:     100.00,
		CreatedBy: adminId,
	}

	mockModelWashPackage := &models.Wash{
		Name:      mockWashPackage.Name,
		Category:  mockWashPackage.Category,
		Price:     float64(mockWashPackage.Price),
		CreatedBy: mockWashPackage.CreatedBy,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	mockRepository.Mock.On("CreateWashPackage", mockModelWashPackage).Return(nil).Run(func(args mock.Arguments) {
		washPackage := args.Get(0).(*models.Wash)
		washPackage.ID = washPackageId
	})

	mockRepository.Mock.On("FindWashPackageByID", washPackageId).Return(&pb.WashPackageData{
		Id:       washPackageId,
		Name:     mockWashPackage.Name,
		Category: mockWashPackage.Category,
		Price:    float32(mockWashPackage.Price),
	}, nil)

	_, err := washStationController.CreateWashPackage(context.Background(), pbRequest)

	assert.Nil(t, err)
	assert.Equal(t, mockWashPackage.Name, pbRequest.Name)
}

func TestFindAllDetailingPackages(t *testing.T) {
	mockDetailingPackage := models.Detailing{
		Name:        "Wash Package 1",
		Description: "Detailing Package 1 Description",
		Price:       100.00,
	}

	mockRepository.Mock.On("FindAllDetailingPackages").Return([]models.Detailing{mockDetailingPackage}, nil)

	_, err := washStationController.FindAllDetailingPackages(context.Background(), &emptypb.Empty{})

	assert.Nil(t, err)
	assert.NotEmpty(t, mockDetailingPackage)
}

func TestCreateDetailingPackage(t *testing.T) {
	var (
		adminId            uint32 = 1
		detailingPackageId uint32 = 1
	)

	pbRequest := &pb.NewDetailingPackageData{
		Name:        "Detailing Package 1",
		Description: "Detailing Package 1 Description",
		Price:       200.00,
		CreatedBy:   adminId,
	}

	mockModelDetailingPackage := &models.Detailing{
		Name:        mockDetailingPackage.Name,
		Description: mockDetailingPackage.Description,
		Price:       float64(mockDetailingPackage.Price),
		CreatedBy:   mockDetailingPackage.CreatedBy,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	mockRepository.Mock.On("CreateDetailingPackage", mockModelDetailingPackage).Return(nil).Run(func(args mock.Arguments) {
		detailingPackage := args.Get(0).(*models.Detailing)
		detailingPackage.ID = detailingPackageId
	})

	mockRepository.Mock.On("FindWashPackageByID", detailingPackageId).Return(&pb.DetailingPackageData{
		Id:          detailingPackageId,
		Name:        mockDetailingPackage.Name,
		Description: mockDetailingPackage.Description,
		Price:       float32(mockDetailingPackage.Price),
	}, nil)

	_, err := washStationController.CreateDetailingPackage(context.Background(), pbRequest)

	assert.Nil(t, err)
	assert.Equal(t, mockDetailingPackage.Name, pbRequest.Name)
}
