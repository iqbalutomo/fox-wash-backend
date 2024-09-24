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
	//detailingpackagemethods
	CreateDetailingPackage(data *models.Detailing) error
	FindAllDetailingPackages() ([]models.Detailing, error)
	FindDetailingPackageByID(DetailingPackageID uint32) (dto.DetailingPackageDataCompact, error)
	UpdateDetailingPackage(DetailingPackageID uint32, data *pb.UpdateDetailingPackageData) error
	DeleteDetailingPackage(DetailingPackageID uint32) error
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

//detailing package repo

func (w *WashStationRepository) CreateDetailingPackage(data *models.Detailing) error {
	if err := w.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (w *WashStationRepository) FindAllDetailingPackages() ([]models.Detailing, error) {
	var detailing []models.Detailing
	if err := w.db.Find(&detailing).Error; err != nil {
		return nil, err
	}
	return detailing, nil
}

func (w *WashStationRepository) FindDetailingPackageByID(DetailingPackageID uint32) (dto.DetailingPackageDataCompact, error) {
	var detailing models.Detailing
	if err := w.db.First(&detailing, DetailingPackageID).Error; err != nil {
		return dto.DetailingPackageDataCompact{}, err
	}
	return dto.DetailingPackageDataCompact{
		ID:          detailing.ID,
		Name:        detailing.Name,
		Description: detailing.Description,
		Price:       detailing.Price,
	}, nil
}

func (w *WashStationRepository) UpdateDetailingPackage(DetailingPackageID uint32, data *pb.UpdateDetailingPackageData) error {
	if err := w.db.Model(&models.Detailing{}).Where("id = ?", DetailingPackageID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (w *WashStationRepository) DeleteDetailingPackage(DetailingPackageID uint32) error {
	if err := w.db.Delete(&models.Detailing{}, DetailingPackageID).Error; err != nil {
		return err
	}
	return nil
}
