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

type PaymentSuccessData struct {
	InvoiceID   string `json:"invoice_id"`
	Status      string `json:"status"`
	Method      string `json:"method"`
	CompletedAt string `json:"completed_at"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}
