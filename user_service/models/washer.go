package models

type Washer struct {
	UserID         uint `gorm:"primaryKey"`
	IsActive       bool `gorm:"not null"`
	WasherStatusID uint `gorm:"not null"`
}

type WasherStatus struct {
	ID      uint   `gorm:"primaryKey"`
	Status  string `gorm:"not null"`
	Washers []Washer
}
