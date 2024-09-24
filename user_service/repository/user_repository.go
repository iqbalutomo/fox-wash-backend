package repository

import (
	"errors"
	"user_service/dto"
	"user_service/models"
	"user_service/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(data *models.User) error
	AddToken(data *models.EmailVerification) error
	VerifyNewUser(id uint32, token string) error
	GetUser(email string) (dto.UserJoinedData, error)
	CreateWasher(userID uint32) error
	WasherActivation(email string) error
	GetWasher(userID int32) (dto.WasherData, error)
	SetWasherStatusOnline(userID uint32) error
	GetAvailableWasher() (dto.WasherOrderData, error)
	SetWasherStatusWashing(washerID int32) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &UserRepository{db}
}

func (u *UserRepository) CreateUser(data *models.User) error {
	result := u.db.Create(data)
	if err := result.Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return status.Error(codes.AlreadyExists, err.Error())
		}

		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.Internal, result.Error.Error())
	}

	return nil
}

func (u *UserRepository) AddToken(data *models.EmailVerification) error {
	if err := u.db.Create(data).Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (u *UserRepository) VerifyNewUser(userID uint32, token string) error {
	verificationData := models.EmailVerification{
		UserID: uint(userID),
	}

	result := u.db.Model(&verificationData).Where("token = ?", token).Update("is_verified", true)
	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.Unauthenticated, "invalid verification credentials")
	}

	return nil
}

func (u *UserRepository) GetUser(email string) (dto.UserJoinedData, error) {
	var userData dto.UserJoinedData

	result := u.db.Table("users u").
		Select("u.id, u.first_name, u.last_name, u.email, u.password, r.name AS role, u.created_at, v.is_verified").
		Where("u.email = ?", email).
		Joins("JOIN roles r ON u.role_id = r.id").
		Joins("JOIN email_verifications v ON v.user_id = u.id").
		Take(&userData)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserJoinedData{}, status.Error(codes.NotFound, err.Error())
		}

		return userData, status.Error(codes.Internal, err.Error())
	}

	return userData, nil
}

func (u *UserRepository) CreateWasher(userID uint32) error {
	washerData := models.Washer{
		UserID:         uint(userID),
		IsOnline:       false,
		WasherStatusID: utils.InActiveWasherStatusID,
		IsActive:       false,
	}

	result := u.db.Create(&washerData)
	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (u *UserRepository) WasherActivation(email string) error {
	var user models.User

	if err := u.db.Preload("Washer").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if user.Washer.UserID == 0 {
		return errors.New("user does not have a washer profile")
	}

	user.Washer.IsActive = true

	if err := u.db.Save(&user.Washer).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetWasher(userID int32) (dto.WasherData, error) {
	var washerData dto.WasherData

	result := u.db.Table("washers").
		Where("user_id = ?", userID).
		Take(&washerData)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.WasherData{}, status.Error(codes.NotFound, err.Error())
		}

		return washerData, status.Error(codes.Internal, err.Error())
	}

	return washerData, nil
}

func (u *UserRepository) SetWasherStatusOnline(userID uint32) error {
	result := u.db.Table("washers").
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"is_online":        true,
			"washer_status_id": 1,
		})
	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.InvalidArgument, "Invalid washer ID")
	}

	return nil
}

func (u *UserRepository) GetAvailableWasher() (dto.WasherOrderData, error) {
	var washerData dto.WasherOrderData

	result := u.db.Table("washers w").
		Select("w.user_id AS id, u.first_name || ' ' || u.last_name AS name, ws.status").
		Joins("JOIN users u on w.user_id = u.id").
		Joins("JOIN washer_statuses ws on w.washer_status_id = ws.id").
		Where("ws.status = ? AND w.is_online = ?", "available", true).
		Take(&washerData)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.WasherOrderData{}, status.Error(codes.Unavailable, "washers not available to take order")
		}

		return dto.WasherOrderData{}, status.Error(codes.Internal, err.Error())
	}

	if err := u.SetWasherStatusWashing(int32(washerData.ID)); err != nil {
		return dto.WasherOrderData{}, err
	}

	return washerData, nil
}

func (u *UserRepository) SetWasherStatusWashing(washerID int32) error {
	result := u.db.Table("washers").
		Where("user_id = ?", washerID).
		Update("washer_status_id", utils.WashingWasherStatusID)
	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if result.RowsAffected == 0 {
		return status.Error(codes.InvalidArgument, "Invalid washer ID")
	}

	return nil
}
