package models

type Order struct {
	OrderDetail OrderDetail `json:"order_detail"`
	User        User        `json:"user"`
	Washer      Washer      `json:"washer"`
	Address     Address     `json:"address"`
	Payment     Payment     `json:"payment"`
	Status      string      `json:"status"`
}

type OrderDetail struct {
	WashPackage      []WashPackage      `json:"wash_packages"`
	DetailingPackage []DetailingPackage `json:"detailing_packages"`
	AppFee           string             `json:"app_fee"`
	TotalPrice       string             `json:"total_price"`
}

type WashPackage struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category int    `json:"category"`
	Price    string `json:"price"`
	Qty      uint   `json:"qty"`
	SubTotal string `json:"subtotal"`
}

type DetailingPackage struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Qty         uint   `json:"qty"`
	SubTotal    string `json:"subtotal"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Washer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Address struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Payment struct {
	InvoiceID  string  `json:"invoice_id"`
	InvoiceURL string  `json:"invoice_url"`
	Total      float32 `json:"total"`
	Method     string  `json:"method"`
	Status     string  `json:"status"`
}

type PaymentSuccess struct {
	InvoiceID   string `json:"invoice_id"`
	Status      string `json:"status"`
	Method      string `json:"method"`
	CompletedAt string `json:"completed_at"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}
