package models

import "time"

type Transaction struct {
	Id             int       `json:"id"`
	DeliveryMethod string    `json:"string"`
	FullName       string    `json:"fullName"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	SubTotal       int       `json:"subTotal"`
	Tax            int       `json:"tax"`
	Total          int       `json:"total"`
	Date           time.Time `json:"date"`
	Status         string    `json:"status"`
	PaymentMethod  string    `json:"paymentMethod"`
}

type TransactionProduct struct {
	Id            int `json:"id"`
	ProductId     int `json:"productId"`
	TransactionId int `json:"transactionId"`
	Qty           int `json:"qty"`
	SizeId        int `json:"sizeId"`
	VariantId     int `json:"variantId"`
	Price         int `json:"price"`
}
