package models

type User struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	FirstName         string `json:"first_name" gorm:"not null"`
	LastName          string `json:"last_name" gorm:"not null"`
	Email             string `json:"email" gorm:"not null; unique"`
	Password          string `json:"password" gorm:"not null"`
	RoleID            uint32 `json:"role_id" gorm:"not null"`
	CreatedAt         string `json:"created_at" gorm:"not null;type:timestamp"`
	EmailVerification EmailVerification
	Washer            Washer
}

type Role struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Users []User
}

type EmailVerification struct {
	UserID     uint   `json:"user_id" gorm:"primaryKey"`
	Token      string `json:"token" gorm:"not null"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
}
