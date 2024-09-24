package repository

import (
	"errors"
	"wash_station_service/dto"
	"wash_station_service/models"
	"wash_station_service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := w.db.Create(data).Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (w *WashStationRepository) FindAllWashPackages() ([]models.Wash, error) {
	var washPackages []models.Wash

	if err := w.db.Find(&washPackages).Error; err != nil {
		return nil, err
	}

	return washPackages, nil
}

func (w *WashStationRepository) FindWashPackageByID(WashPackageID uint32) (dto.WashPackageDataCompact, error) {
	var washPackage dto.WashPackageDataCompact

	if err := w.db.Table("washes").Where("id = ?", WashPackageID).Take(&washPackage).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.WashPackageDataCompact{}, status.Error(codes.NotFound, err.Error())
		}

		return dto.WashPackageDataCompact{}, status.Error(codes.Internal, err.Error())
	}

	return washPackage, nil
}

func (w *WashStationRepository) UpdateWashPackage(WashPackageID uint32, data *pb.UpdateWashPackageData) error {
	washPackage := models.Wash{ID: WashPackageID}

	res := w.db.Model(&washPackage).Updates(models.Wash{
		Name:     data.Name,
		Category: data.Category,
		Price:    float64(data.Price),
	})
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}

		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (w *WashStationRepository) DeleteWashPackage(WashPackageID uint32) error {
	washPackageData := models.Wash{ID: WashPackageID}

	res := w.db.Delete(&washPackageData, "id = ?", WashPackageID)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}

		return status.Error(codes.Internal, err.Error())
	}

	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "invalid wash package ID")
	}

	return nil
}

//detailing package repo

func (w *WashStationRepository) CreateDetailingPackage(data *models.Detailing) error {
	if err := w.db.Create(data).Error; err != nil {
		return status.Error(codes.Internal, err.Error())
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.DetailingPackageDataCompact{}, status.Error(codes.NotFound, err.Error())
		}

		return dto.DetailingPackageDataCompact{}, status.Error(codes.Internal, err.Error())
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}

		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (w *WashStationRepository) DeleteDetailingPackage(DetailingPackageID uint32) error {
	if err := w.db.Delete(&models.Detailing{}, DetailingPackageID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}

		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
