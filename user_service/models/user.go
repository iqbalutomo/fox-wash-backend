package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null; unique"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"not null;type:timestamp"`
}
