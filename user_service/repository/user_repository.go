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
		Select("u.id, u.first_name, u.last_name, u.email, u.password, u.created_at, v.is_verified").
		Where("u.email = ?", email).
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
		IsActive:       false,
		WasherStatusID: utils.InActiveWasherStatusID,
	}

	result := u.db.Create(&washerData)
	if err := result.Error; err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
