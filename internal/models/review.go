package models

type Review struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	Message string `json:"message"`
	Rating  int    `json:"rating"`
}
