package dto

type UserJoinedData struct {
	ID         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Role       string `json:"role_id"`
	CreatedAt  string `json:"created_at"`
	IsVerified bool   `json:"is_verified"`
}
