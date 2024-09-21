package repository

import (
	"user_service/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(*models.User) error
	AddToken(*models.EmailVerification) error
	VerifyNewUser(id uint32, token string) error
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
