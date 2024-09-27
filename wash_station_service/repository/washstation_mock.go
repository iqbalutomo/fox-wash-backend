package repository

import (
	"wash_station_service/dto"
	"wash_station_service/models"
	"wash_station_service/pb"

	"github.com/stretchr/testify/mock"
)

type MockWashStationRepository struct {
	Mock mock.Mock
}

func NewMockWashStationRepository() MockWashStationRepository {
	return MockWashStationRepository{}
}

func (m *MockWashStationRepository) FindAllWashPackages() ([]models.Wash, error) {
	args := m.Mock.Called()
	return args.Get(0).([]models.Wash), args.Error(1)
}

func (m *MockWashStationRepository) FindWashPackageByID(id uint32) (dto.WashPackageDataCompact, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(dto.WashPackageDataCompact), args.Error(1)
}

func (m *MockWashStationRepository) FindMultipleWashPackages(washPackagesId []uint32) ([]*pb.WashPackageCompact, error) {
	args := m.Mock.Called(washPackagesId)
	return args.Get(0).([]*pb.WashPackageCompact), args.Error(1)
}

func (m *MockWashStationRepository) CreateWashPackage(wash *models.Wash) error {
	args := m.Mock.Called(wash)
	return args.Error(0)
}

func (m *MockWashStationRepository) UpdateWashPackage(washPackageId uint32, data *pb.UpdateWashPackageData) error {
	args := m.Mock.Called(washPackageId, data)
	return args.Error(0)
}

func (m *MockWashStationRepository) DeleteWashPackage(washPackageId uint32) error {
	args := m.Mock.Called(washPackageId)
	return args.Error(0)
}

func (m *MockWashStationRepository) CreateDetailingPackage(detailing *models.Detailing) error {
	args := m.Mock.Called(detailing)
	return args.Error(0)
}

func (m *MockWashStationRepository) FindAllDetailingPackages() ([]models.Detailing, error) {
	args := m.Mock.Called()
	return args.Get(0).([]models.Detailing), args.Error(1)
}

func (m *MockWashStationRepository) FindDetailingPackageByID(id uint32) (dto.DetailingPackageDataCompact, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(dto.DetailingPackageDataCompact), args.Error(1)
}

func (m *MockWashStationRepository) FindMultipleDetailingPackages(detailingPackagesId []uint32) ([]*pb.DetailingPackageCompact, error) {
	args := m.Mock.Called(detailingPackagesId)
	return args.Get(0).([]*pb.DetailingPackageCompact), args.Error(1)
}

func (m *MockWashStationRepository) UpdateDetailingPackage(detailingPackageId uint32, data *pb.UpdateDetailingPackageData) error {
	args := m.Mock.Called(detailingPackageId, data)
	return args.Error(0)
}

func (m *MockWashStationRepository) DeleteDetailingPackage(detailingPackageId uint32) error {
	args := m.Mock.Called(detailingPackageId)
	return args.Error(0)
}
