package models

type Cart struct {
	Id        int `json:"id"`
	UserId    int `json:"userId"`
	ProductId int `json:"productId"`
}
