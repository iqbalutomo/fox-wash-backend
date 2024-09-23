package repository

import (
	"errors"
	"wash_station_service/dto"
	"wash_station_service/models"
	"wash_station_service/pb"

	"gorm.io/gorm"
)

type WashStation interface {
	CreateWashPackage(data *models.Wash) error
	FindAllWashPackages() ([]models.Wash, error)
	FindWashPackageByID(WashPackageID uint32) (dto.WashPackageDataCompact, error)
	UpdateWashPackage(WashPackageID uint32, data *pb.UpdateWashPackageData) error
	DeleteWashPackage(WashPackageID uint32) error
}

type WashStationRepository struct {
	db *gorm.DB
}

func NewWashStationRepository(db *gorm.DB) WashStation {
	return &WashStationRepository{db}
}

func (w *WashStationRepository) CreateWashPackage(data *models.Wash) error {
	return errors.New("") // TODO: logic here
}

func (w *WashStationRepository) FindAllWashPackages() ([]models.Wash, error) {
	return nil, errors.New("") // TODO: logic here
}

func (w *WashStationRepository) FindWashPackageByID(WashPackageID uint32) (dto.WashPackageDataCompact, error) {
	return dto.WashPackageDataCompact{}, errors.New("") // TODO: logic here
}

func (w *WashStationRepository) UpdateWashPackage(WashPackageID uint32, data *pb.UpdateWashPackageData) error {
	return errors.New("") // TODO: logic here
}

func (w *WashStationRepository) DeleteWashPackage(WashPackageID uint32) error {
	return errors.New("") // TODO: logic here
}
