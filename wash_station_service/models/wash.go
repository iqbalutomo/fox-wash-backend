package models

type Wash struct {
	ID        uint32  `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"not null"`
	Category  uint32  `json:"category" gorm:"not null"` // cc bike
	Price     float64 `json:"price" gorm:"not null"`
	CreatedBy uint32  `json:"created_by" gorm:"not null"`
	CreatedAt string  `json:"created_at" gorm:"not null;type:timestamp"`
}
