package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderDetail OrderDetail        `bson:"order_detail" json:"order_detail"`
	User        User               `bson:"user" json:"user"`
	Washer      Washer             `bson:"washer" json:"washer"`
	Address     Address            `bson:"address" json:"address"`
	Payment     Payment            `bson:"payment" json:"payment"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
}

type OrderDetail struct {
	WashPackage      []WashPackage      `bson:"wash_packages" json:"wash_packages"`
	DetailingPackage []DetailingPackage `bson:"detailing_packages" json:"detailing_packages"`
	AppFee           float32            `bson:"app_fee" json:"app_fee"`
	TotalPrice       float32            `bson:"total_price" json:"total_price"`
}

type WashPackage struct {
	ID       int     `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Category int     `bson:"category" json:"category"`
	Price    float32 `bson:"price" json:"price"`
	Qty      uint    `bson:"qty" json:"qty"`
}

type DetailingPackage struct {
	ID          int     `bson:"id" json:"id"`
	Name        string  `bson:"name" json:"name"`
	Description string  `bson:"description" json:"description"`
	Price       float32 `bson:"price" json:"price"`
	Qty         uint    `bson:"qty" json:"qty"`
}

type User struct {
	ID    int    `bson:"id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

type Washer struct {
	ID     int    `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Status string `bson:"status" json:"status"`
}

type Address struct {
	Latitude  float32 `bson:"latitude" json:"latitude"`
	Longitude float32 `bson:"longitude" json:"longitude"`
}

type Payment struct {
	InvoiceID  string  `bson:"invoice_id" json:"invoice_id"`
	InvoiceURL string  `bson:"invoice_url" json:"invoice_url"`
	Total      float32 `bson:"total" json:"total"`
	Method     string  `bson:"method" json:"method"`
	Status     string  `bson:"status" json:"status"`
}
