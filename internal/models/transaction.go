package models

import "time"

type Transaction struct {
	Id             int64             `json:"id" db:"id"`
	UserId         string            `json:"user_id" db:"user_id"`
	TrxCode        string            `json:"trx_code" db:"trx_code"`
	DeliveryMethod string            `json:"delivery_method" db:"delivery_method"`
	FullName       string            `json:"full_name" db:"full_name"`
	Email          string            `json:"email" db:"email"`
	Address        string            `json:"address" db:"address"`
	SubTotal       int               `json:"sub_total" db:"sub_total"`
	Tax            int               `json:"tax" db:"tax"`
	Total          int               `json:"total" db:"total"`
	Date           time.Time         `json:"date" db:"date"`
	Status         string            `json:"status" db:"status"`
	PaymentMethod  string            `json:"payment_method" db:"payment_method"`
	Items          []TransactionItem `json:"items,omitzero"`
}

type TransactionItem struct {
	Id            int    `json:"id" db:"id"`
	ProductId     int    `json:"product_id" db:"product_id"`
	TransactionId int64  `json:"transaction_id" db:"transaction_id"`
	Quantity      int    `json:"quantity" db:"quantity"`
	SizeId        *int   `json:"size_id" db:"size_id"`
	VariantId     *int   `json:"variant_id" db:"variant_id"`
	Price         int    `json:"price" db:"price"`
	ProductName   string `json:"product_name" db:"product_name"`
	Image         string `json:"image" db:"image"`
	SizeName      string `json:"size_name" db:"size_name"`
	VariantName   string `json:"variant_name" db:"variant_name"`
}
