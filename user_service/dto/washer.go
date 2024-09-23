package dto

type WasherData struct {
	UserID         uint32 `json:"user_id"`
	IsOnline       bool   `json:"is_online"`
	WasherStatusID uint32 `json:"washer_status_id"`
	IsActive       bool   `json:"is_active"`
}
