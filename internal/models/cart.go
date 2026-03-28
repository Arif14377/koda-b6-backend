package models

import "time"

type Cart struct {
	Id        int       `json:"id"`
	UserId    string    `json:"userId"`
	ProductId int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	SizeId    *int      `json:"sizeId,omitzero"`
	VariantId *int      `json:"variantId,omitzero"`
	CreatedAt time.Time `json:"createdAt"`
}

type CartItemResponse struct {
	Id          int     `json:"id" db:"id"`
	ProductId   int     `json:"productId" db:"product_id"`
	ProductName string  `json:"productName" db:"name"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       int     `json:"price" db:"price"`
	Image       string  `json:"image" db:"image"`
	Size        *string `json:"size,omitzero" db:"size"`
	Variant     *string `json:"variant,omitzero" db:"variant"`
	SizeId      *int    `json:"sizeId,omitzero" db:"size_id"`
	VariantId   *int    `json:"variantId,omitzero" db:"variant_id"`
	Total       int     `json:"total" db:"total"`
}
